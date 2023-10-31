

export const noDataProps = {
  headers: [
    {
      displayName: 'test-column-display',
      name: 'test-column-name',
    },
    {
      displayName: 'test-column-display2',
      name: 'test-column-name2',
    },
  ],
  data: [],
  emptyMessage: 'No data to show',
}

export const clickableProps = {
  headers: [
    {
      displayName: 'test-column-display',
      name: 'test-column-name',
    },
    {
      displayName: 'test-column-display2',
      name: 'test-column-name2',
    },
  ],
  data: [
    {
      'test-column-name': 'entry1.1',
      'test-column-name2': 'entry1.2',
      clickable: true,
    },
    {
      'test-column-name': 'entry2.1',
      'test-column-name2': 'entry2.2',
      clickable: true,
    },
  ],
  emptyMessage: 'No data to show',
}

export const paginatedProps = {
  headers: [
    {
      displayName: 'test-column-display',
      name: 'test-column-name',
    },
    {
      displayName: 'test-column-display2',
      name: 'test-column-name2',
    },
  ],
  data: [
    {
      'test-column-name': 'entry1.1',
      'test-column-name2': 'entry1.2',
    },
    {
      'test-column-name': 'entry2.1',
      'test-column-name2': 'entry2.2',
    },
  ],
  paginated: true,
  totalCount: 2,
  pageSize: 2,
  initialPage: 0,
  emptyMessage: 'No data to show',
}

export const sortableProps = {
  headers: [
    {
      displayName: 'test-column-display',
      name: 'test-column-name',
      sortable: true,
    },
    {
      displayName: 'test-column-display2',
      name: 'test-column-name2',
    },
    {
      displayName: 'test-column-display3',
      name: 'test-column-name3',
      sortable: true,
    },
  ],
  data: [
    {
      'test-column-name': 'entry1.1',
      'test-column-name2': 'entry1.2',
      'test-column-name3': 'entry1.3',
    },
    {
      'test-column-name': 'entry2.1',
      'test-column-name2': 'entry2.2',
      'test-column-name3': 'entry2.3',
    },
  ],
  emptyMessage: 'No data to show',
  order: undefined,
  orderBy: undefined,
}
