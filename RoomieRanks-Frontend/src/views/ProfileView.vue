<template>
   <div class="tabs">
        <button 
            class="tab" 
            :class="{ active: activeTab === 'profile' }"
            @click="activeTab = 'profile'">
            User Profile</button>
        <button 
            class="tab"
            :class="{ active: activeTab === 'chores' }"
            @click="activeTab = 'chores'"
        >
        Current Tasks </button>
        <button
            class="tab"
            :class="{ active: activeTab === 'collectibles' }"
            @click="activeTab = 'collectibles'"
        >
        Current Collectibles </button>

        <button
            class="tab"
            :class="{ active: activeTab === 'history' }"
            @click="activeTab = 'history'"
        >
        Transaction History </button>
    </div>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}

.tabs {
    display: flex;
    justify-content: space-between;
    margin-bottom: 1rem;
    margin-top: 100px;
    margin-left: auto;
    margin-right: auto;
    width: 500px;
    background-color: #1e1e1e;
    border-radius: 10px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.tab {
    margin-left: auto;
    margin-right: auto;
    padding: 1rem 2rem;
    cursor: pointer;
    border: none;
    border-radius: 10px;
    background-color: #1e1e1e;
    color: rgba(243, 244, 246, 1);
    transition: background-color 0.3s;
    flex: 1;
    text-align: center;
}

.tab:last-child {
    border-right: none;
}

.tab:hover {
    background: rgba(55, 65, 81, 1);
}

.tab.active {
    background: rgba(55, 65, 81, 1);
    font-weight: bold;
}

.tab-content {
    display: none;
    width: 500px;
    text-align: center;
    margin-top: 1rem;
    background-color: #1e1e1e;
    padding: 1rem;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.tab-content.active {
    display: block;
}
</style>

<script>
export default {
  name: "ProfileView",
  data() {
    return {
      activeTab: "profile"
     }
    }
   } // Default tab
</script>