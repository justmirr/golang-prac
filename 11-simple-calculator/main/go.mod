module main/calculator

go 1.24.3

replace addition => ../addition

replace subtraction => ../subtraction

replace division => ../division

replace multiplication => ../multiplication

require (
	addition v0.0.0-00010101000000-000000000000
	division v0.0.0-00010101000000-000000000000
	multiplication v0.0.0-00010101000000-000000000000
	subtraction v0.0.0-00010101000000-000000000000
)
