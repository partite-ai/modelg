package typeutil

import "go/types"

func InferTypeFromFirstParam(sig *types.Signature, candidateArg1Type types.Type) (*types.Signature, bool) {
	if sig.Params().Len() == 0 {
		return nil, false
	}
	sigArg1Type := sig.Params().At(0).Type()
	typeParams := make([]*types.TypeParam, sig.TypeParams().Len())
	for i := 0; i < sig.TypeParams().Len(); i++ {
		typeParams[i] = sig.TypeParams().At(i)
	}
	u := newUnifier(typeParams, nil)
	if !u.unify(sigArg1Type, candidateArg1Type, exact) {
		return nil, false
	}
	inferredTypes := u.inferred(typeParams)
	for i, t := range inferredTypes {
		if typeParams[i].Constraint() != nil {
			if !types.AssignableTo(t, typeParams[i].Constraint()) {
				return nil, false
			}
		}
		if t == nil {
			return nil, false
		}
	}

	instantiated, err := types.Instantiate(nil, sig, inferredTypes, false)
	if err != nil {
		return nil, false
	}

	instantiatedSig, ok := instantiated.(*types.Signature)
	if !ok {
		return nil, false
	}

	return instantiatedSig, true
}
