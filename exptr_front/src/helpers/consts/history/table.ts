import { date, type QTableColumn } from "quasar";

export const TABLE_COLUMNS: QTableColumn[] = [
  {
    name: 'name',
    required: true,
    label: 'Name',
    align: 'left',
    field: 'name',
    sortable: true
  },
  {
    name: 'type',
    label: 'Type',
    align: 'left',
    field: 'type',
    sortable: true
  },
  {
    name: 'amount',
    label: 'Amount',
    align: 'left',
    field: (row) => row.amount + row.currency,
    sortable: true
  },
  {
    name: 'comment',
    label: 'Comment',
    align: 'left',
    field: 'comment',
  },
  {
    name: 'date',
    label: 'Date',
    align: 'left',
    field: row => date.formatDate(row.updated_at, 'YYYY-MM-DD'),
  },
]