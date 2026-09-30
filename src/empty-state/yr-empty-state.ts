// <yr-empty-state><svg slot="icon">...</svg>No results found.</yr-empty-state>
// — an icon + message placeholder. Light DOM (the "slot" is just a
// [slot="icon"] attribute selector in CSS, not a real shadow-DOM slot); no
// behavior needed, registered only so Storybook has something to catalog.
export class YrEmptyState extends HTMLElement {}

customElements.define('yr-empty-state', YrEmptyState);
