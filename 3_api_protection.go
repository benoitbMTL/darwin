package main

import (
	"github.com/labstack/echo/v4"
)

type PetstorePet []struct {
	ID        int      `json:"id,omitempty"`
	Category  Category `json:"category,omitempty"`
	Name      string   `json:"name,omitempty"`
	PhotoUrls []string `json:"photoUrls,omitempty"`
	Tags      []Tags   `json:"tags,omitempty"`
	Status    string   `json:"status,omitempty"`
}

type Category struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Tags struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func handlePetstoreAPIRequestGet(c echo.Context) error { return()
}

func handlePetstoreAPIRequestPost(c echo.Context) error {return()
}

func handlePetstoreAPIRequestPut(c echo.Context) error {return()
}

func handlePetstoreAPIRequestDelete(c echo.Context) error {return()
}



