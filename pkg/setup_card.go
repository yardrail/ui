package ui

import "github.com/a-h/templ"

// SetupCardProps configures a SetupCard.
type SetupCardProps struct {
	Attrs templ.Attributes
}

// SetupCard renders a setup/onboarding card. Children are the card content.
func SetupCard(p *SetupCardProps) templ.Component {
	return setupCard(p)
}

// SetupCardProgressProps configures a SetupCardProgress.
type SetupCardProgressProps struct {
	Steps   string
	Current string
	Attrs   templ.Attributes
}

// SetupCardProgress renders a step progress indicator.
func SetupCardProgress(p *SetupCardProgressProps) templ.Component {
	return setupCardProgress(p)
}

// SetupCardHeaderProps configures a SetupCardHeader.
type SetupCardHeaderProps struct {
	Attrs templ.Attributes
}

// SetupCardHeader renders the header of a setup card.
func SetupCardHeader(p *SetupCardHeaderProps) templ.Component {
	return setupCardHeader(p)
}

// SetupCardBodyProps configures a SetupCardBody.
type SetupCardBodyProps struct {
	Attrs templ.Attributes
}

// SetupCardBody renders the body of a setup card.
func SetupCardBody(p *SetupCardBodyProps) templ.Component {
	return setupCardBody(p)
}

// SetupCardFooterProps configures a SetupCardFooter.
type SetupCardFooterProps struct {
	Attrs templ.Attributes
}

// SetupCardFooter renders the footer of a setup card.
func SetupCardFooter(p *SetupCardFooterProps) templ.Component {
	return setupCardFooter(p)
}
