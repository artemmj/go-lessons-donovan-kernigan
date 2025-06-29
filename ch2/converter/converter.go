package converter

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string { return fmt.Sprintf("%g meters", m) }
func (f Foot) String() string  { return fmt.Sprintf("%g foots", f) }

func MetersToFoots(m Meter) Foot { return Foot(m * 3.2808) }
func FootsToMeters(f Foot) Meter { return Meter(f / 3.2808) }
