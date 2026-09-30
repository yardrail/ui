import { YrDropdownMenu } from '../dropdown-menu/yr-dropdown-menu';

// <yr-account-menu name="Jordan Ellery" email="jordan@ellerylogistics.com">
//   <span slot="avatar">JE</span>
//   <a class="yr-dropdown-menu-item" href="/account">Account settings</a>
//   ...
// </yr-account-menu>
//
// A higher-level wrapper around yr-dropdown-menu: auto-builds the trigger
// (small avatar + chevron) and panel header (large avatar + name + email)
// from a flat attribute contract, instead of every service hand-assembling
// the identical trigger/header markup + exact class names independently.
// `[slot="avatar"]` supplies the avatar's light-DOM content (initials text
// or an <img>) — cloned into both the small trigger avatar and the large
// header avatar. An optional `[slot="trigger-extra"]` child renders inside
// the trigger button before the avatar (e.g. yardrail-ui's inline
// name/role text) for services that want more than depot's compact
// avatar-only trigger. Everything else passed as children (links, a
// divider) renders below the header, inside the panel. Reuses
// yr-dropdown-menu's open/close/outside-click/Escape behavior via
// inheritance — only the initial DOM restructuring here is new.
export class YrAccountMenu extends YrDropdownMenu {
  connectedCallback() {
    this.restructure();
    super.connectedCallback();
  }

  private restructure() {
    if (this.querySelector(':scope > [slot="trigger"]')) return;

    const name = this.getAttribute('name') ?? '';
    const email = this.getAttribute('email') ?? '';
    const avatarSlot = this.querySelector('[slot="avatar"]');
    const triggerExtra = this.querySelector('[slot="trigger-extra"]');
    const items = Array.from(this.childNodes).filter((node) => node !== avatarSlot && node !== triggerExtra);

    const trigger = document.createElement('button');
    trigger.setAttribute('slot', 'trigger');
    trigger.className = 'yr-account-menu-trigger';
    if (triggerExtra) trigger.append(triggerExtra.cloneNode(true));
    const triggerAvatar = document.createElement('yr-avatar');
    if (avatarSlot) triggerAvatar.append(avatarSlot.cloneNode(true));
    trigger.append(triggerAvatar, this.buildChevron());

    const header = document.createElement('div');
    header.className = 'yr-dropdown-menu-header';
    const headerAvatar = document.createElement('yr-avatar');
    headerAvatar.setAttribute('size', 'lg');
    if (avatarSlot) headerAvatar.append(avatarSlot.cloneNode(true));
    header.append(headerAvatar);

    if (name || email) {
      const headerText = document.createElement('div');
      if (name) {
        const nameEl = document.createElement('div');
        nameEl.className = 'yr-dropdown-menu-header-name';
        nameEl.textContent = name;
        headerText.append(nameEl);
      }
      if (email) {
        const emailEl = document.createElement('div');
        emailEl.className = 'yr-dropdown-menu-header-email';
        emailEl.textContent = email;
        headerText.append(emailEl);
      }
      header.append(headerText);
    }

    const panel = document.createElement('div');
    panel.setAttribute('slot', 'panel');
    panel.append(header, ...items);

    this.replaceChildren(trigger, panel);
  }

  private buildChevron(): SVGSVGElement {
    const svg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
    svg.setAttribute('class', 'yr-dropdown-menu-chevron');
    svg.setAttribute('width', '12');
    svg.setAttribute('height', '12');
    svg.setAttribute('viewBox', '0 0 24 24');
    svg.setAttribute('fill', 'none');
    svg.setAttribute('stroke', 'currentColor');
    svg.setAttribute('stroke-width', '2.5');
    svg.setAttribute('stroke-linecap', 'round');
    svg.setAttribute('stroke-linejoin', 'round');
    const poly = document.createElementNS('http://www.w3.org/2000/svg', 'polyline');
    poly.setAttribute('points', '6 9 12 15 18 9');
    svg.append(poly);
    return svg;
  }
}

customElements.define('yr-account-menu', YrAccountMenu);
