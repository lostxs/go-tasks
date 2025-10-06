package main

import (
	"3-struct/api"
	"3-struct/config"
	"os"
	"testing"
)

var testApi *api.Api

func TestMain(m *testing.M) {
	cfg := config.NewConfig()
	testApi = api.NewApi(cfg)
	os.Exit(m.Run())
}

func createTestBin(t *testing.T) string {
	resp, err := testApi.CreateBin("my.json", "test-bin")
	if err != nil {
		t.Fatalf("Ошибка при создании bin: %v", err)
	}
	return resp.Metadata.ID
}

func deleteTestBin(t *testing.T, id string) {
	_, err := testApi.DeleteBin(id)
	if err != nil {
		t.Errorf("Не удалось удалить bin %s: %v", id, err)
	}
}

func TestCreateBin(t *testing.T) {
	id := createTestBin(t)
	if id == "" {
		t.Fatal("Создание bin вернуло пустой ID")
	}
	deleteTestBin(t, id)
}

func TestUpdateBin(t *testing.T) {
	id := createTestBin(t)

	_, err := testApi.UpdateBin("my.json", id)
	if err != nil {
		t.Fatalf("Ошибка при обновлении bin: %v", err)
	}

	resp, err := testApi.GetBin(id)
	if err != nil {
		t.Fatalf("Ошибка при получении обновлённого bin: %v", err)
	}
	if resp.Metadata.ID != id {
		t.Errorf("Ожидали ID %s, получили %s", id, resp.Metadata.ID)
	}

	deleteTestBin(t, id)
}

func TestGetBin(t *testing.T) {
	id := createTestBin(t)

	resp, err := testApi.GetBin(id)
	if err != nil {
		t.Fatalf("Ошибка при получении bin: %v", err)
	}
	if resp.Metadata.ID != id {
		t.Errorf("Ожидали ID %s, получили %s", id, resp.Metadata.ID)
	}

	deleteTestBin(t, id)
}

func TestDeleteBin(t *testing.T) {
	id := createTestBin(t)

	_, err := testApi.DeleteBin(id)
	if err != nil {
		t.Fatalf("Ошибка при удалении bin: %v", err)
	}

	_, err = testApi.GetBin(id)
	if err == nil {
		t.Error("Bin должен быть удалён, но всё ещё доступен")
	}
}
