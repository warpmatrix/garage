func isNumber(s string) bool {
    intReg := regexp.MustCompile(`[+-]?\d+`)
    decReg := regexp.MustCompile(`[+-]?(\d*\.\d+|\d+\.\d*)`)
    expReg := regexp.MustCompile(`[eE]` + intReg.String())
    numReg := regexp.MustCompile(`^((` + intReg.String() + `)|(` +
        decReg.String() + `))` + `(` + expReg.String() + `)?$`)
    return numReg.Find([]byte(s)) != nil
}
