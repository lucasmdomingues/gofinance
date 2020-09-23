# Installation

```go 
go get github.com/lucasmdomingues/gofinance
```

# Example

### Get only quotes for currencies and exchanges.

```go
import (
	"fmt"
	"log"
	"github.com/lucasmdomingues/gofinance"
)

func main() {
	service := gofinance.NewService("KEY")

	quotations, err := service.GetQuotations()
	if err != nil {
		log.Fatal(err)
	}
}
```

### Obtain only CDI and SELIC rates.

```go
import (
	"fmt"
	"log"
	"github.com/lucasmdomingues/gofinance"
)

func main() {
	service := gofinance.NewService("KEY")

	taxes, err := service.GetTaxes()
	if err != nil {
		log.Fatal(err)
	}
}

```
### HG Finance
https://hgbrasil.com/status/finance/

### TO DO

* Obtendo dados históricos
* Personalizando a resposta
