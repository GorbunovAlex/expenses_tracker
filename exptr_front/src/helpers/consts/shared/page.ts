import { PageAction, type IPageAction, type IPageLink } from "@/helpers/types";

export const PAGE_LINKS: IPageLink[] = [
  {
    icon: "euro_symbol",
    to: "/",
  },
  {
    icon: "analytics",
    to: "/analytics",
  },
  {
    icon: "schedule",
    to: "/history",
  }
];

export const PAGE_ACTIONS: IPageAction[] = [
  {
    label: "Logout",
    action: PageAction.LOGOUT
  }    
]