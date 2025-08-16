<template>
  <div id="profile">
    <!-- Loading State -->
    <div v-if="loading" class="loading-container">
      <div class="loading-spinner"></div>
      <p>Loading profile...</p>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="error-container">
      <h2>Error Loading Profile</h2>
      <p>{{ error }}</p>
      <button @click="loadUserProfile" class="retry-button">Retry</button>
    </div>

    <!-- Profile Content -->
    <main v-else class="profile-content">
      <!-- User Information Section -->
      <section class="user-info">
        <h2 class="section-title">User Information</h2>
        <div class="info-card">
          <p><strong>Username:</strong> {{ userData.username }}</p>
          <p><strong>Email:</strong> {{ userData.email }}</p>
          <p><strong>Name:</strong> {{ userData.name }}</p>
        </div>
      </section>

      <!-- XP and Level Section -->
      <section class="xp-info">
        <h2 class="section-title">Progress</h2>
        <div class="info-card">
          <p><strong>Level:</strong> {{ userData.level }}</p>
          <p><strong>Title:</strong> {{ userData.title }}</p>
          <div class="progress-bar-container">
            <div class="progress-bar" :style="{ width: xpProgress + '%' }"></div>
          </div>
          <p class="xp-text">XP: {{ userData.xp }} / {{ xpForNextLevel }}</p>
        </div>
      </section>

             <!-- User Collection Section -->
       <section class="user-collection">
         <h2 class="section-title">Your Collection</h2>
         <div class="collection-container">
           <div v-for="card in collection" :key="card.id" class="card">
             <img :src="card.img" :alt="card.name" class="card-img" />
             <div class="card-details">
               <p class="card-name">{{ card.name }}</p>
               <p class="card-rarity">Rarity: {{ card.rarity }}</p>
               <p class="card-type">Type: {{ card.type }}</p>
             </div>
           </div>
         </div>
       </section>

       <!-- Debug Section -->
       <section class="debug-section">
         <h2 class="section-title">Debug Information</h2>
         <div class="info-card">
           <p><strong>Session Username:</strong> {{ session.username }}</p>
           <p><strong>Session Permission Level:</strong> {{ session.permissionLevel }}</p>
           <p><strong>SessionStorage Username:</strong> {{ sessionStorage.getItem('loggedInUsername') }}</p>
           <p><strong>SessionStorage Permission:</strong> {{ sessionStorage.getItem('permissionLevel') }}</p>
           <button @click="refreshSession" class="debug-button">Refresh Session</button>
           <button @click="testWithUsername" class="debug-button">Test with 'timber'</button>
         </div>
       </section>
    </main>
  </div>
</template>

<script>
import { fetchUser } from '@/utils/api';
import { session } from '@/session';

export default {
  name: "UserProfilePage",
  data() {
    return {
      loading: true,
      error: null,
      userData: {
        username: '',
        name: '',
        email: '',
        household_id: 0,
        title: '',
        level: 1,
        xp: 0
      },
      collection: [
        {
          id: 1,
          img: "https://via.placeholder.com/100x150",
          name: "Golden Spoon",
          rarity: "Rare",
          type: "Tool",
        },
        {
          id: 2,
          img: "https://via.placeholder.com/100x150",
          name: "Mighty Mop",
          rarity: "Epic",
          type: "Weapon",
        },
        {
          id: 3,
          img: "https://via.placeholder.com/100x150",
          name: "Dishwasher Pro",
          rarity: "Legendary",
          type: "Appliance",
        },
      ],
    };
  },
  computed: {
    xpProgress() {
      const xpForCurrentLevel = this.getXPForLevel(this.userData.level);
      const xpForNextLevel = this.getXPForLevel(this.userData.level + 1);
      const xpInCurrentLevel = this.userData.xp - xpForCurrentLevel;
      const xpNeededForNextLevel = xpForNextLevel - xpForCurrentLevel;
      
      if (xpNeededForNextLevel <= 0) return 100;
      return Math.min(100, (xpInCurrentLevel / xpNeededForNextLevel) * 100);
    },
    xpForNextLevel() {
      return this.getXPForLevel(this.userData.level + 1);
    }
  },
  methods: {
         async loadUserProfile() {
       this.loading = true;
       this.error = null;
       
       console.log('🚀 Starting to load user profile...');
       
       // Force refresh session from storage in case there's a timing issue
       session.refreshFromStorage();
       
       console.log('👤 Current session username:', session.username);
       console.log('👤 Session permission level:', session.permissionLevel);
       console.log('🔍 SessionStorage loggedInUsername:', sessionStorage.getItem('loggedInUsername'));
       
       try {
         // Use the logged-in username from session, or fallback to a default
         const username = session.username !== 'guest' ? session.username : 'testuser';
         console.log('🎯 Attempting to fetch user with username:', username);
         
         if (username === 'undefined' || username === undefined) {
           throw new Error('Username is undefined. Please log in again.');
         }
         
         const userData = await fetchUser(username);
         console.log('✅ User data received:', userData);
         this.userData = userData;
       } catch (err) {
         console.error('💥 Error in loadUserProfile:', err);
         this.error = err.message || 'Failed to load user profile';
       } finally {
         this.loading = false;
                  console.log('🏁 Finished loading attempt');
       }
     },
     refreshSession() {
       console.log('🔄 Manual session refresh requested');
       session.refreshFromStorage();
       this.loadUserProfile();
     },
     async testWithUsername() {
       console.log('🧪 Testing with hardcoded username "timber"');
       try {
         const userData = await fetchUser('timber');
         console.log('✅ Test successful:', userData);
         this.userData = userData;
         this.error = null;
       } catch (err) {
         console.error('❌ Test failed:', err);
         this.error = err.message;
       }
     },
     getXPForLevel(level) {
      // Simple XP calculation: each level requires level * 100 XP
      // Level 1: 0 XP, Level 2: 100 XP, Level 3: 300 XP, etc.
      return (level - 1) * 100;
    }
  },
  mounted() {
    this.loadUserProfile();
  }
};
</script>

<style scoped>
/* General Reset */
#profile {
  margin: 0;
  padding: 5rem 0;
  font-family: 'Roboto', sans-serif;
  color: #fff;
  min-height: 100vh;
  padding-bottom: 4rem;
}

/* Profile Header */
.profile-header {
  text-align: center;
  padding: 2rem;
  background-color: #1e40af;
  color: #fca400;
  border-radius: 10px;
  margin: 0 auto 2rem;
  width: 90%;
  max-width: 800px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
  position: relative;
  top: -1rem;
}

.profile-title {
  font-size: 2.8rem;
  font-weight: bold;
  text-transform: uppercase;
  margin: 0;
}

/* Profile Content */
.profile-content {
  max-width: 800px;
  margin: 0 auto;
  padding: 1rem 1.5rem;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 15px;
  box-shadow: 0 6px 15px rgba(0, 0, 0, 0.3);
}

.section-title {
  font-size: 1.85rem;
  margin-bottom: 1rem;
  color: #fca400;
  text-align: center;
  font-weight: 600;
}

.info-card {
  background-color: rgba(255, 255, 255, 0.1);
  border-radius: 10px;
  padding: 1.8rem;
  margin-bottom: 2rem;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.info-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.3);
}

.info-card p {
  font-size: 1.2rem;
  margin: 0.8rem 0;
  color: #e5e7eb;
}

/* Progress Bar */
.progress-bar-container {
  background-color: rgba(255, 255, 255, 0.2);
  border-radius: 20px;
  overflow: hidden;
  margin: 1.2rem 0;
  height: 28px;
  box-shadow: inset 0 2px 5px rgba(0, 0, 0, 0.3);
}

.progress-bar {
  height: 100%;
  background: linear-gradient(90deg, #fca400, #ffcd3c);
  transition: width 0.4s ease;
}

.xp-text {
  text-align: center;
  font-size: 1.25rem;
  margin-top: 0.8rem;
  color: #e5e7eb;
}

/* Collection Cards */
.collection-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 1.5rem;
  margin-top: 1.5rem;
}

.card {
  background-color: rgba(255, 255, 255, 0.1);
  border-radius: 10px;
  padding: 1rem;
  text-align: center;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.card:hover {
  transform: translateY(-5px);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.3);
}

.card-img {
  max-width: 100%;
  border-radius: 5px;
  margin-bottom: 0.5rem;
}

.card-details {
  color: #e5e7eb;
}

.card-name {
  font-size: 1.2rem;
  font-weight: bold;
  margin-bottom: 0.5rem;
}

 .card-rarity,
 .card-type {
   font-size: 1rem;
   margin: 0.3rem 0;
 }

 /* Loading and Error States */
 .loading-container,
 .error-container {
   display: flex;
   flex-direction: column;
   align-items: center;
   justify-content: center;
   min-height: 50vh;
   text-align: center;
 }

 .loading-spinner {
   width: 50px;
   height: 50px;
   border: 4px solid rgba(255, 255, 255, 0.3);
   border-top: 4px solid #fca400;
   border-radius: 50%;
   animation: spin 1s linear infinite;
   margin-bottom: 1rem;
 }

 @keyframes spin {
   0% { transform: rotate(0deg); }
   100% { transform: rotate(360deg); }
 }

 .error-container h2 {
   color: #ef4444;
   margin-bottom: 1rem;
 }

 .error-container p {
   color: #e5e7eb;
   margin-bottom: 1.5rem;
 }

 .retry-button {
   background-color: #fca400;
   color: #1e40af;
   border: none;
   padding: 0.75rem 1.5rem;
   border-radius: 8px;
   font-size: 1rem;
   font-weight: 600;
   cursor: pointer;
   transition: background-color 0.3s ease;
 }

 .retry-button:hover {
   background-color: #ffcd3c;
 }

 .debug-button {
   background-color: #3b82f6;
   color: white;
   border: none;
   padding: 0.75rem 1.5rem;
   border-radius: 8px;
   font-size: 1rem;
   font-weight: 600;
   cursor: pointer;
   transition: background-color 0.3s ease;
   margin: 0.5rem;
 }

 .debug-button:hover {
   background-color: #2563eb;
 }
 </style>
  