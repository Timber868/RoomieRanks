import { reactive } from 'vue';

export const session = reactive({
  username: sessionStorage.getItem('loggedInUsername') || 'guest',
  permissionLevel: sessionStorage.getItem('permissionLevel') ? Number(sessionStorage.getItem('permissionLevel')) : 0,

  updateSession(username: string, permissionLevel: number) {
    this.username = username;
    this.permissionLevel = permissionLevel;
    sessionStorage.setItem('loggedInUsername', username);
    sessionStorage.setItem('permissionLevel', String(permissionLevel)); // Convert to string before saving
  },

  logout() {
    this.username = 'guest';
    this.permissionLevel = 0;
    sessionStorage.setItem('loggedInUsername', 'guest');
    sessionStorage.setItem('permissionLevel', '0'); // Ensure it's a string
  }
});
