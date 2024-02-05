export interface IPageLink {
  icon: string;
  to: string;
}

export interface IPageAction {
  label: string
  action: PageAction;
}

export enum PageAction {
  LOGOUT = "logout",
}