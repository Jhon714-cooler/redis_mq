package mq

type ProducerOptions struct{
	MsgQuenueLen	int

}
type ProducerOption func(p *ProducerOptions)  
	
func WithMsgQueueLen(MsgQuenueLen int) ProducerOption {
	return func (opts *ProducerOptions) {
		opts.MsgQuenueLen = MsgQuenueLen
	}
}
