// <yr-workflow-card> — an installed workflow summary card. Light-DOM
// children include a header (name + status pill), stats row, and footer
// (last-run text + trigger button). Purely presentational, no behavior.
export class YrWorkflowCard extends HTMLElement {}

customElements.define('yr-workflow-card', YrWorkflowCard);
