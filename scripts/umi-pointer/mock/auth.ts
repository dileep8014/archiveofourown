export default {
  'GET /auth': (req: any, res: any) => {
    res.setHeader('x-auth-token', 'token-token-token');
    res.send('');
  },

};
