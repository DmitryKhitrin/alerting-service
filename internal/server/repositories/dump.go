package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	files, err := ioutil.ReadDir("/tmp/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}

	file, err := os.Open(l.cfg.FileName)
	if err != nil {
		log.Println("error while reading file", err)
		return err
	}

	defer func() {
		err := file.Close()
		log.Println("file closed")
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

	log.Println("save", l.cfg.StoreInterval, l.cfg.FileName)
	log.Println("interval", l.cfg.StoreInterval.Seconds(), l.cfg.FileName)

	if l.cfg.StoreInterval.Seconds() != 0 && l.cfg.FileName != "" {
		log.Println("start saving")
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
		defer stop()

		metricSaver := func(storeInterval time.Duration) {
			ticker := time.NewTicker(storeInterval)

			for {
				select {
				case <-ticker.C:
					err := l.SaveToFile()
					if err != nil {
						log.Println("not saved in loop", err)
					}
					log.Println("saved")
				case <-ctx.Done():
					ticker.Stop()
					return
				}
			}
		}

		go metricSaver(l.cfg.StoreInterval)
	}
}
