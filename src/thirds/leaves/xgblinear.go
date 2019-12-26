package leaves

// xgLinear is XGBoost model (gblinear)
type xgLinear struct {
	NumFeature int
	nClasses   int
	BaseScore  float64
	Weights    []float32
}

func (e *xgLinear) NEstimators() int {
	return 1
}

func (e *xgLinear) NClasses() int {
	return e.nClasses
}

func (e *xgLinear) adjustNEstimators(nEstimators int) int {
	// gbliearn has only one estimator per class
	return 1
}

func (e *xgLinear) NFeatures() int {
	return e.NumFeature
}

func (e *xgLinear) Name() string {
	return "xgboost.gblinear"
}

func (e *xgLinear) predictInner(fvals []float64, nIterations int, predictions []float64, startIndex int) {
	for k := 0; k < e.nClasses; k++ {
		predictions[startIndex+k] = e.BaseScore + float64(e.Weights[e.nClasses*e.NumFeature+k])
		for i := 0; i < e.NumFeature; i++ {
			predictions[startIndex+k] += fvals[i] * float64(e.Weights[e.nClasses*i+k])
		}
	}
}

func (e *xgLinear) resetFVals(fvals []float64) {
	for j := 0; j < len(fvals); j++ {
		fvals[j] = 0.0
	}
}
