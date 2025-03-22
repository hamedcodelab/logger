package customCore

import (
	"fmt"
	"go.uber.org/zap/zapcore"
)

// NewCore creates a Core that writes logs to a WriteSyncer.
func NewCore(enc zapcore.Encoder, ws zapcore.WriteSyncer, enab zapcore.LevelEnabler) zapcore.Core {
	return &ioCore{
		LevelEnabler: enab,
		enc:          enc,
		out:          ws,
	}
}

type ioCore struct {
	zapcore.LevelEnabler
	enc zapcore.Encoder
	out zapcore.WriteSyncer
}

func (c *ioCore) Level() zapcore.Level {
	return zapcore.LevelOf(c.LevelEnabler)
}

func addFields(enc zapcore.ObjectEncoder, fields []zapcore.Field) {
	for i := range fields {
		fields[i].AddTo(enc)
	}
}

func (c *ioCore) With(fields []zapcore.Field) zapcore.Core {
	clone := c.clone()
	addFields(clone.enc, fields)
	return clone
}

func (c *ioCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *ioCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	buf, err := c.enc.EncodeEntry(ent, fields)
	if err != nil {
		return err
	}

	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", getZapLevelColor(ent.Level), buf.String())
	_, err = c.out.Write([]byte(colored))
	buf.Free()
	if err != nil {
		return err
	}
	if ent.Level > zapcore.ErrorLevel {
		// Since we may be crashing the program, sync the output.
		// Ignore Sync errors, pending a clean solution to issue #370.
		_ = c.Sync()
	}
	return nil
}

func (c *ioCore) Sync() error {
	return c.out.Sync()
}

func (c *ioCore) clone() *ioCore {
	return &ioCore{
		LevelEnabler: c.LevelEnabler,
		enc:          c.enc.Clone(),
		out:          c.out,
	}
}
