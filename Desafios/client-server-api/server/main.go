package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// CoinDetails guarda as informações de uma moeda específica
type CoinDetails struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

// ExchangeCoin agora aceita qualquer par de moedas dinamicamente
type ExchangeCoin map[string]CoinDetails

// MoedaDB é o seu modelo do BANCO DE DADOS (GORM)
type CoinDB struct {
	ID         uint `gorm:"primaryKey"`
	Code       string
	Codein     string
	Name       string
	High       string
	Low        string
	VarBid     string
	PctChange  string
	Bid        string
	Ask        string
	Timestamp  string
	CreateDate string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type Env struct {
	DB *gorm.DB
}

func main() {
	// Initialize DB
	dsn := "root:root@tcp(localhost:3306)/client-server-api"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print("Erro ao estabelcer conexão com banco de dados")
		return
	}

	db.AutoMigrate(&CoinDB{})

	log.Print("Conexão com DB estabelecida!")

	sqlDB, err := db.DB()
	if err != nil {
		log.Print(err)
	}

	defer sqlDB.Close()

	env := &Env{DB: db}

	http.HandleFunc("/cotacao", env.handler)
	http.ListenAndServe(":8080", nil)
}

func (e *Env) saveCoin(c *CoinDB, ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancel()

	exchangeCoin := c
	e.DB.WithContext(ctx).Create(&exchangeCoin)
}

func (e *Env) handler(w http.ResponseWriter, r *http.Request) {
	coin := r.URL.Query().Get("moeda")
	exchangeURL := fmt.Sprintf("https://economia.awesomeapi.com.br/json/last/%s", coin)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, exchangeURL, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprint(w, "O servidor demorou muito para responder. Por favor, tente novamente mais tarde.")
			log.Print("O servidor de exchange demorou muito a responder")
			return
		}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var data ExchangeCoin
	err = json.Unmarshal(body, &data)

	// 2. Pega os dados da moeda específica de dentro do mapa (ex: "USDBRL")
	clearCoin := strings.ReplaceAll(coin, "-", "")
	coinData, ok := data[clearCoin]

	if !ok {
		log.Println("Moeda não encontrada no JSON")
		return
	}

	// 3. Converte os dados da API para o formato que o GORM aceita (MoedaDB)
	coinToSave := CoinDB{
		Code:       coinData.Code,
		Codein:     coinData.Codein,
		Name:       coinData.Name,
		High:       coinData.High,
		Low:        coinData.Low,
		VarBid:     coinData.VarBid,
		PctChange:  coinData.PctChange,
		Bid:        coinData.Bid,
		Ask:        coinData.Ask,
		Timestamp:  coinData.Timestamp,
		CreateDate: coinData.CreateDate,
	}

	w.Write(body)
	e.saveCoin(&coinToSave, r.Context())
}
