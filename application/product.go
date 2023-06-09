package application

type ProductInterface interface {
	IsValid() (bool, error)
	Enabled() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	status string
}

func (p *Product) IsValid() (bool, error) {

}

func (p *Product) Enabled() error {

}

func (p *Product) Disable() error {

}

func (p *Product) GetID() string {

}

func (p *Product) GetName() string {

}

func (p *Product) GetStatus() string {

}

func (p *Product) GetPrice() float64 {

}
