package uuid5

import (
    "github.com/satori/go.uuid"
    "errors"
    "fmt"
)

type UUIDGenerator struct {
    Namespace string
}

func (u UUIDGenerator) GenerateWithName(name string) (string, error) {

    ns, err := uuid.FromString(u.Namespace)
    if err != nil {
        return "", errors.New(fmt.Sprintf("[namespace] is invalid (%s)", err.Error()))
    }

    if name == "" {
        return "", errors.New("[name] should not be empty")
    }

    generated := uuid.NewV5(ns, name)
    return generated.String(), nil
}

func (u UUIDGenerator) IsValid(ident string) error {

    _, err := uuid.FromString(ident)
    if err != nil {
        return err
    }

    return nil
}
