import umiRequest from 'umi-request';

export async function QueryCurrentUser() {
  return umiRequest('/api/v1/users/currentUser');
}
