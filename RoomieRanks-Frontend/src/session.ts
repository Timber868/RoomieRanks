import { reactive } from 'vue';

export const session = reactive({
  username: sessionStorage.getItem('loggedInUsername') || 'guest',
  permissionLevel: sessionStorage.getItem('permissionLevel') ? Number(sessionStorage.getItem('permissionLevel')) : 0,

  updateSession(username: string, permissionLevel: number) {
    console.log('🔄 Updating session with:', { username, permissionLevel });
    this.username = username;
    this.permissionLevel = permissionLevel;
    sessionStorage.setItem('loggedInUsername', username);
    sessionStorage.setItem('permissionLevel', String(permissionLevel)); // Convert to string before saving
    console.log('✅ Session updated. Current state:', { 
      username: this.username, 
      permissionLevel: this.permissionLevel,
      sessionStorageUsername: sessionStorage.getItem('loggedInUsername'),
      sessionStoragePermission: sessionStorage.getItem('permissionLevel')
    });
  },

  logout() {
    console.log('🚪 Logging out...');
    this.username = 'guest';
    this.permissionLevel = 0;
    sessionStorage.setItem('loggedInUsername', 'guest');
    sessionStorage.setItem('permissionLevel', '0'); // Ensure it's a string
    console.log('✅ Logged out. Session state:', { username: this.username, permissionLevel: this.permissionLevel });
  },

  refreshFromStorage() {
    console.log('🔄 Refreshing session from storage...');
    const storedUsername = sessionStorage.getItem('loggedInUsername');
    const storedPermission = sessionStorage.getItem('permissionLevel');
    
    console.log('📦 Stored values:', { storedUsername, storedPermission });
    
    this.username = storedUsername || 'guest';
    this.permissionLevel = storedPermission ? Number(storedPermission) : 0;
    
    console.log('✅ Session refreshed:', { username: this.username, permissionLevel: this.permissionLevel });
  }
});

// Debug initial session state
console.log('🎬 Session initialized with:', { 
  username: session.username, 
  permissionLevel: session.permissionLevel,
  sessionStorageUsername: sessionStorage.getItem('loggedInUsername'),
  sessionStoragePermission: sessionStorage.getItem('permissionLevel')
});
