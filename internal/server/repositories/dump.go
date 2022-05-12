package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func (l *LocalStorageRepository) TryRestore() error {
	if !l.cfg.ShouldRestore || l.cfg.FileName == "" {
		return nil
	}

	file, err := os.OpenFile(l.cfg.FileName, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println("error while reading file", err)
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Println("error while closing file", err)
		}
	}()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&l.repository)
	if err != nil {
		log.Println("error while decoding file", err)
		return err
	}

	return nil
}

func (l *LocalStorageRepository) SaveToFile() error {
	if l.cfg.FileName == "" {
		log.Println("trying to store metrics with empty file")
		return nil
	}

	file, err := os.OpenFile(l.cfg.FileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	encoder := json.NewEncoder(file)

	l.mutex.RLock()
	err = encoder.Encode(l.repository)
	if err != nil {
		log.Println("error while saving file")
	}
	l.mutex.RUnlock()
	return err
}

func (l *LocalStorageRepository) RunDataDumper(ctx context.Context) {
	if l.cfg.StoreInterval.Seconds() != 0 && l.cfg.FileName != "" {

		metricSaver := func(storeInterval time.Duration) {
			ticker := time.NewTicker(storeInterval)

			for {
				select {
				case <-ticker.C:
					err := l.SaveToFile()
					if err != nil {
						log.Println("error while saving by interval", err)
					}
				case <-ctx.Done():
					ticker.Stop()
					return
				}
			}
		}

		go metricSaver(l.cfg.StoreInterval)
	}
}
