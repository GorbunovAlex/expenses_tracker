module.exports = {
  'exptr': {
    output: {
      mode: 'single',
      target: './src/helpers/api/api.ts',
      schemas: './src/helpers/api/models',
      baseUrl: '/swagger/doc.json',
      override: {
        mutator: {
          path: './src/helpers/api/axios-instance.ts',
          name: 'instance',
        },
      },
    },
    input: {
      target: 'http://localhost:3000/swagger/doc.json',
    }
  },
};