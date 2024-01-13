package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// Agents lists agents and their status.
func Agents(ctx context.Context, s AMISocket) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionAgents)
	callback := NewAmiCallbackService(ctx, s, c, []string{config.AmiListenerEventAgents},
		[]string{config.AmiListenerEventAgentsComplete})
	return callback.SendSuperLevel()
}

// AgentLogoff sets an agent as no longer logged in.
func AgentLogoff(ctx context.Context, s AMISocket, agent string, soft bool) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionAgentLogOff)
	c.SetV(map[string]interface{}{
		config.AmiFieldAgent: agent,
		config.AmiFieldSoft:  soft,
	})
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}
