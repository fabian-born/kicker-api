package main

import (
        "database/sql"
        b64 "encoding/base64"
        "encoding/json"
        "log"
        "net/http"
        "strings"

        "github.com/gin-gonic/gin"
        _ "github.com/go-sql-driver/mysql"
)


// team functions //
