package services

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/NeginSal/api-hub/internal/config"
    "github.com/NeginSal/api-hub/internal/models"
)

func FetchAndStoreLMSData() error {
    resp, err := http.Get("http://localhost:3000/lms_mock.json")
    if err != nil {
        return fmt.Errorf("failed to fetch data: %v", err)
    }
    defer resp.Body.Close()

    var data []models.LMSData
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return fmt.Errorf("failed to decode response: %v", err)
    }

    collection := config.GetCollection("apihub", "lms")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var docs []interface{}
    for _, d := range data {
        docs = append(docs, d)
    }

    _, err = collection.InsertMany(ctx, docs)
    if err != nil {
        return fmt.Errorf("failed to insert data: %v", err)
    }

    return nil
}
