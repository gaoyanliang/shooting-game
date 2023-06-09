package main

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
)

type Config struct {
	ScreenWidth        int        `json:"screenWidth"`
	ScreenHeight       int        `json:"screenHeight"`
	Title              string     `json:"title"`
	BgColor            color.RGBA `json:"bgColor"`
	ShipSpeedFactor    float64    `json:"shipSpeedFactor"`
	BulletWidth        int        `json:"bulletWidth"`
	BulletHeight       int        `json:"bulletHeight"`
	BulletSpeedFactor  float64    `json:"bulletSpeedFactor"`
	BulletColor        color.RGBA `json:"bulletColor"`
	MaxBulletNum       int        `json:"maxBulletNum"`
	BulletInterval     int64      `json:"bulletInterval"`
	MonsterSpeedFactor float64    `json:"monsterSpeedFactor"`
	FontSize           int        `json:"fontSize"`
	TitleFontSize      int        `json:"titleFontSize"`
	SmallFontSize      int        `json:"smallFontSize"`
	FailCount          int        `json:"failCount"`
}

func loadConfig() *Config {
	f, err := os.Open("./config.json")
	if err != nil {
		log.Fatalf("[loadConfig] os.Open failed: %v\n", err)
	}

	var cfg Config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatalf("[loadConfig] json.Decode failed: %v\n", err)
	}

	return &cfg
}
