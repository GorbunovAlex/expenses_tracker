import { PageAction } from "@/helpers/types";
import { Dialog } from "quasar";

import CategoryDialog from "@/components/categories/dialogs/CategoryDialog.vue";

export default function usePage() {
  function actionHandler<T>(action: PageAction, payload: T) {
    switch (action) {
      case PageAction.LOGOUT:
        console.log("Logout");
        break;
      case PageAction.TOGGLE_CATEGORY:
        Dialog.create({
          component: CategoryDialog,
          componentProps: {
            id: payload
          }
        })
        break;
    }
  }

  return {
    actionHandler
  }
}