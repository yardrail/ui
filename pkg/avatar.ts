// <yr-avatar size="lg">{{ initials }}</yr-avatar> — a circular initials
// badge. Content (initials) is plain light-DOM text supplied by the caller.
// Light DOM, styled by yr-avatar.css; purely presentational, no behavior.
export class YrAvatar extends HTMLElement {
  static observedAttributes = ['size'];
}

customElements.define('yr-avatar', YrAvatar);
