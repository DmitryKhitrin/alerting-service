package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (l *LocalStorageRepository) TryRestore() error {
	if !l.cfg.ShouldRestore || l.cfg.FileName == "" {
		return nil
	}

	file, err := os.Open(l.cfg.FileName)
	if err != nil {
		log.Println("error while reading file")
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&l.repository)
	if err != nil {
		log.Println("error while decoding file")
		return err
	}

	return nil
}

func (l *LocalStorageRepository) SaveToFile() error {
	if l.cfg.FileName == "" {
		fmt.Println("trying to store metrics with empty file")
		return nil
	}

	file, err := os.Create(l.cfg.FileName)
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
	l.mutex.RUnlock()
	return err
}

func (l *LocalStorageRepository) RunDataDumper() {

	if l.cfg.StoreInterval.Seconds() != 0 && l.cfg.FileName != "" {

		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
		defer stop()

		metricSaver := func(storeInterval time.Duration) {
			ticker := time.NewTicker(storeInterval)

			for {
				select {
				case <-ticker.C:
					err := l.SaveToFile()
					if err != nil {
						log.Println(err)
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
