package main

import (
    "fmt"
    "flag"
    "regexp"
    "os"
)

func main() {
    var (
        main_order_id = flag.String("m", "", "main_order_id")
        azoya_order_id = flag.String("a", "", "azoya order id")
    )

    flag.Parse()

    val, _ := regexp.MatchString("^[0-9]+$", *main_order_id)
    if !val {
        fmt.Println("-m need number.")
        os.Exit(1)
    }

    fmt.Printf("INSERT INTO orders_external VALUES(0, '%s', '%s', 'AZOYA', now());\n", *main_order_id, *azoya_order_id)
    fmt.Println("")
    fmt.Printf("UPDATE customer_orders SET status = 1 WHERE external_order_id IN ('%s');\n", *azoya_order_id)
    fmt.Printf("UPDATE customer_orders SET synchronized_status = 0 WHERE external_order_id IN ('%s');\n", *azoya_order_id)
}
