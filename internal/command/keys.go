package command

import "github.com/nalgeon/redka"

// Returns all key names that match a pattern.
// KEYS pattern
// https://redis.io/commands/keys
type Keys struct {
	baseCmd
	pattern string
}

func parseKeys(b baseCmd) (*Keys, error) {
	cmd := &Keys{baseCmd: b}
	if len(cmd.args) != 1 {
		return cmd, ErrInvalidArgNum
	}
	cmd.pattern = string(cmd.args[0])
	return cmd, nil
}

func (cmd *Keys) Run(w Writer, red *redka.Tx) (any, error) {
	keys, err := red.Key().Keys(cmd.pattern)
	if err != nil {
		w.WriteError(cmd.Error(err))
		return nil, err
	}
	w.WriteArray(len(keys))
	for _, key := range keys {
		w.WriteBulkString(key.Key)
	}
	return keys, nil
}
