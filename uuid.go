package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func uuid1() {
	fmt.Println("========== uuid1() ==========")
	id := uuid.New()

	fmt.Println("id", id)
	fmt.Println("id.Domain", id.Domain())
	fmt.Println("id.ID", id.ID())
	fmt.Println("id.Time", id.Time())
	fmt.Println("id.String", id.String())

	fmt.Println("time.Now()", time.Now())
	fmt.Println("time.Now().Unix()", time.Now().Unix())
}
