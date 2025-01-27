<script setup>
import { RouterLink, RouterView } from 'vue-router';
import { session } from './session.ts';
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

onMounted(() => {
  // Initialize session storage if not set
  if (!sessionStorage.getItem("loggedInUsername")) {
    sessionStorage.setItem("loggedInUsername", "guest");
    session.loggedInUsername = "guest";
  }
  if (!sessionStorage.getItem("permissionLevel")) {
    sessionStorage.setItem("permissionLevel", 0);
    session.permissionLevel = 0;
  }
});

function handleLogout() {
  session.logout();
  router.push("/");
}
</script>

<template>
  <div id="app">
    <header>
      <nav class="navbar">
        <div class="nav-container">
          <div class="nav-logo">
            <span class="logo-text">
              <span class="logo-roomie">Roomie</span>
              <span class="logo-ranks">Ranks</span>
            </span>
          </div>
          <div class="nav-links">
            <RouterLink to="/" class="nav-item">Home</RouterLink> 
            <!-- <RouterLink to="/trade" class="nav-item">Trade</RouterLink> -->
            <!-- <RouterLink to="/money" class="nav-item">Money Management</RouterLink> -->
            <RouterLink to="/household" class="nav-item">No Household</RouterLink>
            <RouterLink to="/hasHousehold" class="nav-item">Yes Household</RouterLink>
            <RouterLink to="/profile" class="nav-item">Profile</RouterLink>
            <RouterLink to="/login" class="nav-item">Login/Register</RouterLink>
            <button class="nav-item" v-if="session.permissionLevel != 0" @click="handleLogout">Logout</button>
          </div>
        </div>
      </nav>
    </header>

    <main>
      <RouterView />
    </main>
  </div>
</template>

<style>
{
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  width: 100vw;
}

header {
  width: 100%;
  background-color: #000000; 
  position: fixed;
  top: 0;
  left: 0;
  z-index: 1000;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.navbar {
  max-width: 1200px; 
  margin: 0 auto;
  padding: 0.5rem 1rem;
  margin-left: 10%; 
}

.nav-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.nav-logo {
  display: flex;
  align-items: center;
  margin-left: -85px;
}

.nav-links {
  display: flex;
  gap: 1rem;
}

.nav-item {
  color: #fff;
  text-decoration: none;
  padding: 10px 15px;
  border-radius: 8px;
  transition: background-color 0.3s ease, transform 0.2s ease;
  background-color: transparent; 
}

.nav-item:hover {
  background-color: rgba(252, 164, 0, 0.2);
  transform: scale(1.05);
}

.nav-item:active {
  background-color: rgba(252, 164, 0, 0.3);
  transform: scale(0.98);
}

.nav-item:focus {
  outline: none;
  box-shadow: 0 0 0 3px rgba(252, 164, 0, 0.5);
}

.router-link-exact-active.nav-item {
  background-color: rgba(252, 164, 0, 0.3);
}

.router-link-active.nav-item {
  background-color: transparent; 
}

main {
  flex: 1;
  padding: 0;
  margin: 0 auto;
  width: 100%;
}

.logo-image {
  height: 35px;
  width: auto;
  display: block;
}

button.nav-item {
  all: unset; 
  color: #fff;
  cursor: pointer;
  padding: 10px 15px;
  border-radius: 8px;
  transition: background-color 0.3s ease, transform 0.2s ease;
  background-color: transparent; 
}

button.nav-item:hover {
  background-color: rgba(252, 164, 0, 0.2);
  transform: scale(1.05);
}

button.nav-item:active {
  background-color: rgba(252, 164, 0, 0.3);
  transform: scale(0.98);
}

button.nav-item:focus {
  outline: none;
  box-shadow: 0 0 0 3px rgba(252, 164, 0, 0.5);
}

.logo-text {
  font-family: 'Roboto', sans-serif; /* Use a boxy font or a similar one */
  font-size: 1.5rem; /* Adjust the size as needed */
  font-weight: bold;
  display: flex;
  align-items: center;
  gap: 5px;
}

.logo-roomie {
  color: #00adef; /* Blue from Portal's color scheme */
  text-transform: uppercase; /* Optional for a boxy effect */
}

.logo-ranks {
  color: #fca400; /* Orange from Portal's color scheme */
  text-transform: uppercase;
}
</style>
