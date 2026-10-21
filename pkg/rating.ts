// <yr-rating> — stars and review count are server-rendered by the templ
// component. This class exists only to register the custom element tag.
export class YrRating extends HTMLElement {}

customElements.define('yr-rating', YrRating);
