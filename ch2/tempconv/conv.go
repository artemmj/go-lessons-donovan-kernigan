package tempconv

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func CtoK(c Celsius) Kelvin     { return Kelvin(c - AbsolteZeroC) }

func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func FtoK(f Fahrenheit) Kelvin  { return Kelvin((f + 459.67) * 5 / 9) }

func KtoC(k Kelvin) Celsius    { return Celsius(k - Kelvin(AbsolteZeroC)) }
func KtoF(k Kelvin) Fahrenheit { return Fahrenheit((k * 9 / 5) - 459.67) }
