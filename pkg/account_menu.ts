import { YrDropdownMenu } from './dropdown_menu';

// <yr-account-menu> — extends yr-dropdown-menu for open/close/Escape
// behavior. All DOM structure is server-rendered by the templ component.
export class YrAccountMenu extends YrDropdownMenu {}

customElements.define('yr-account-menu', YrAccountMenu);
