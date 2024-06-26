// Code generated by counterfeiter. DO NOT EDIT.
package tokenfakes

import (
	"context"
	"sync"

	"platform_engineer_clone/models"
)

type FakeDataPersistence struct {
	GenerateStub        func(context.Context, int, int, int) (string, error)
	generateMutex       sync.RWMutex
	generateArgsForCall []struct {
		arg1 context.Context
		arg2 int
		arg3 int
		arg4 int
	}
	generateReturns struct {
		result1 string
		result2 error
	}
	generateReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetAllStub        func(context.Context) ([]models.Token, error)
	getAllMutex       sync.RWMutex
	getAllArgsForCall []struct {
		arg1 context.Context
	}
	getAllReturns struct {
		result1 []models.Token
		result2 error
	}
	getAllReturnsOnCall map[int]struct {
		result1 []models.Token
		result2 error
	}
	GetTokenStub        func(context.Context, string) (*models.Token, error)
	getTokenMutex       sync.RWMutex
	getTokenArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getTokenReturns struct {
		result1 *models.Token
		result2 error
	}
	getTokenReturnsOnCall map[int]struct {
		result1 *models.Token
		result2 error
	}
	RevokeTokenStub        func(context.Context, string) error
	revokeTokenMutex       sync.RWMutex
	revokeTokenArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	revokeTokenReturns struct {
		result1 error
	}
	revokeTokenReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateTokenToExpiredStub        func(context.Context, *models.Token) error
	updateTokenToExpiredMutex       sync.RWMutex
	updateTokenToExpiredArgsForCall []struct {
		arg1 context.Context
		arg2 *models.Token
	}
	updateTokenToExpiredReturns struct {
		result1 error
	}
	updateTokenToExpiredReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDataPersistence) Generate(arg1 context.Context, arg2 int, arg3 int, arg4 int) (string, error) {
	fake.generateMutex.Lock()
	ret, specificReturn := fake.generateReturnsOnCall[len(fake.generateArgsForCall)]
	fake.generateArgsForCall = append(fake.generateArgsForCall, struct {
		arg1 context.Context
		arg2 int
		arg3 int
		arg4 int
	}{arg1, arg2, arg3, arg4})
	stub := fake.GenerateStub
	fakeReturns := fake.generateReturns
	fake.recordInvocation("Generate", []interface{}{arg1, arg2, arg3, arg4})
	fake.generateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDataPersistence) GenerateCallCount() int {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	return len(fake.generateArgsForCall)
}

func (fake *FakeDataPersistence) GenerateCalls(stub func(context.Context, int, int, int) (string, error)) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = stub
}

func (fake *FakeDataPersistence) GenerateArgsForCall(i int) (context.Context, int, int, int) {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	argsForCall := fake.generateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeDataPersistence) GenerateReturns(result1 string, result2 error) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = nil
	fake.generateReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDataPersistence) GenerateReturnsOnCall(i int, result1 string, result2 error) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = nil
	if fake.generateReturnsOnCall == nil {
		fake.generateReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.generateReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDataPersistence) GetAll(arg1 context.Context) ([]models.Token, error) {
	fake.getAllMutex.Lock()
	ret, specificReturn := fake.getAllReturnsOnCall[len(fake.getAllArgsForCall)]
	fake.getAllArgsForCall = append(fake.getAllArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetAllStub
	fakeReturns := fake.getAllReturns
	fake.recordInvocation("GetAll", []interface{}{arg1})
	fake.getAllMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDataPersistence) GetAllCallCount() int {
	fake.getAllMutex.RLock()
	defer fake.getAllMutex.RUnlock()
	return len(fake.getAllArgsForCall)
}

func (fake *FakeDataPersistence) GetAllCalls(stub func(context.Context) ([]models.Token, error)) {
	fake.getAllMutex.Lock()
	defer fake.getAllMutex.Unlock()
	fake.GetAllStub = stub
}

func (fake *FakeDataPersistence) GetAllArgsForCall(i int) context.Context {
	fake.getAllMutex.RLock()
	defer fake.getAllMutex.RUnlock()
	argsForCall := fake.getAllArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDataPersistence) GetAllReturns(result1 []models.Token, result2 error) {
	fake.getAllMutex.Lock()
	defer fake.getAllMutex.Unlock()
	fake.GetAllStub = nil
	fake.getAllReturns = struct {
		result1 []models.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeDataPersistence) GetAllReturnsOnCall(i int, result1 []models.Token, result2 error) {
	fake.getAllMutex.Lock()
	defer fake.getAllMutex.Unlock()
	fake.GetAllStub = nil
	if fake.getAllReturnsOnCall == nil {
		fake.getAllReturnsOnCall = make(map[int]struct {
			result1 []models.Token
			result2 error
		})
	}
	fake.getAllReturnsOnCall[i] = struct {
		result1 []models.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeDataPersistence) GetToken(arg1 context.Context, arg2 string) (*models.Token, error) {
	fake.getTokenMutex.Lock()
	ret, specificReturn := fake.getTokenReturnsOnCall[len(fake.getTokenArgsForCall)]
	fake.getTokenArgsForCall = append(fake.getTokenArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetTokenStub
	fakeReturns := fake.getTokenReturns
	fake.recordInvocation("GetToken", []interface{}{arg1, arg2})
	fake.getTokenMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDataPersistence) GetTokenCallCount() int {
	fake.getTokenMutex.RLock()
	defer fake.getTokenMutex.RUnlock()
	return len(fake.getTokenArgsForCall)
}

func (fake *FakeDataPersistence) GetTokenCalls(stub func(context.Context, string) (*models.Token, error)) {
	fake.getTokenMutex.Lock()
	defer fake.getTokenMutex.Unlock()
	fake.GetTokenStub = stub
}

func (fake *FakeDataPersistence) GetTokenArgsForCall(i int) (context.Context, string) {
	fake.getTokenMutex.RLock()
	defer fake.getTokenMutex.RUnlock()
	argsForCall := fake.getTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDataPersistence) GetTokenReturns(result1 *models.Token, result2 error) {
	fake.getTokenMutex.Lock()
	defer fake.getTokenMutex.Unlock()
	fake.GetTokenStub = nil
	fake.getTokenReturns = struct {
		result1 *models.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeDataPersistence) GetTokenReturnsOnCall(i int, result1 *models.Token, result2 error) {
	fake.getTokenMutex.Lock()
	defer fake.getTokenMutex.Unlock()
	fake.GetTokenStub = nil
	if fake.getTokenReturnsOnCall == nil {
		fake.getTokenReturnsOnCall = make(map[int]struct {
			result1 *models.Token
			result2 error
		})
	}
	fake.getTokenReturnsOnCall[i] = struct {
		result1 *models.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeDataPersistence) RevokeToken(arg1 context.Context, arg2 string) error {
	fake.revokeTokenMutex.Lock()
	ret, specificReturn := fake.revokeTokenReturnsOnCall[len(fake.revokeTokenArgsForCall)]
	fake.revokeTokenArgsForCall = append(fake.revokeTokenArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.RevokeTokenStub
	fakeReturns := fake.revokeTokenReturns
	fake.recordInvocation("RevokeToken", []interface{}{arg1, arg2})
	fake.revokeTokenMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDataPersistence) RevokeTokenCallCount() int {
	fake.revokeTokenMutex.RLock()
	defer fake.revokeTokenMutex.RUnlock()
	return len(fake.revokeTokenArgsForCall)
}

func (fake *FakeDataPersistence) RevokeTokenCalls(stub func(context.Context, string) error) {
	fake.revokeTokenMutex.Lock()
	defer fake.revokeTokenMutex.Unlock()
	fake.RevokeTokenStub = stub
}

func (fake *FakeDataPersistence) RevokeTokenArgsForCall(i int) (context.Context, string) {
	fake.revokeTokenMutex.RLock()
	defer fake.revokeTokenMutex.RUnlock()
	argsForCall := fake.revokeTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDataPersistence) RevokeTokenReturns(result1 error) {
	fake.revokeTokenMutex.Lock()
	defer fake.revokeTokenMutex.Unlock()
	fake.RevokeTokenStub = nil
	fake.revokeTokenReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDataPersistence) RevokeTokenReturnsOnCall(i int, result1 error) {
	fake.revokeTokenMutex.Lock()
	defer fake.revokeTokenMutex.Unlock()
	fake.RevokeTokenStub = nil
	if fake.revokeTokenReturnsOnCall == nil {
		fake.revokeTokenReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.revokeTokenReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDataPersistence) UpdateTokenToExpired(arg1 context.Context, arg2 *models.Token) error {
	fake.updateTokenToExpiredMutex.Lock()
	ret, specificReturn := fake.updateTokenToExpiredReturnsOnCall[len(fake.updateTokenToExpiredArgsForCall)]
	fake.updateTokenToExpiredArgsForCall = append(fake.updateTokenToExpiredArgsForCall, struct {
		arg1 context.Context
		arg2 *models.Token
	}{arg1, arg2})
	stub := fake.UpdateTokenToExpiredStub
	fakeReturns := fake.updateTokenToExpiredReturns
	fake.recordInvocation("UpdateTokenToExpired", []interface{}{arg1, arg2})
	fake.updateTokenToExpiredMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDataPersistence) UpdateTokenToExpiredCallCount() int {
	fake.updateTokenToExpiredMutex.RLock()
	defer fake.updateTokenToExpiredMutex.RUnlock()
	return len(fake.updateTokenToExpiredArgsForCall)
}

func (fake *FakeDataPersistence) UpdateTokenToExpiredCalls(stub func(context.Context, *models.Token) error) {
	fake.updateTokenToExpiredMutex.Lock()
	defer fake.updateTokenToExpiredMutex.Unlock()
	fake.UpdateTokenToExpiredStub = stub
}

func (fake *FakeDataPersistence) UpdateTokenToExpiredArgsForCall(i int) (context.Context, *models.Token) {
	fake.updateTokenToExpiredMutex.RLock()
	defer fake.updateTokenToExpiredMutex.RUnlock()
	argsForCall := fake.updateTokenToExpiredArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDataPersistence) UpdateTokenToExpiredReturns(result1 error) {
	fake.updateTokenToExpiredMutex.Lock()
	defer fake.updateTokenToExpiredMutex.Unlock()
	fake.UpdateTokenToExpiredStub = nil
	fake.updateTokenToExpiredReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDataPersistence) UpdateTokenToExpiredReturnsOnCall(i int, result1 error) {
	fake.updateTokenToExpiredMutex.Lock()
	defer fake.updateTokenToExpiredMutex.Unlock()
	fake.UpdateTokenToExpiredStub = nil
	if fake.updateTokenToExpiredReturnsOnCall == nil {
		fake.updateTokenToExpiredReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateTokenToExpiredReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDataPersistence) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	fake.getAllMutex.RLock()
	defer fake.getAllMutex.RUnlock()
	fake.getTokenMutex.RLock()
	defer fake.getTokenMutex.RUnlock()
	fake.revokeTokenMutex.RLock()
	defer fake.revokeTokenMutex.RUnlock()
	fake.updateTokenToExpiredMutex.RLock()
	defer fake.updateTokenToExpiredMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDataPersistence) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
