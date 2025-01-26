<template>
    <div id="profile">
  
      <!-- Profile Content -->
      <main class="profile-content">
        <!-- User Information Section -->
        <section class="user-info">
          <h2 class="section-title">User Information</h2>
          <div class="info-card">
            <p><strong>Username:</strong> choreMaster123</p>
            <p><strong>Email:</strong> choremaster@example.com</p>
            <p><strong>Name:</strong> Alex Johnson</p>
          </div>
        </section>
  
        <!-- XP and Level Section -->
        <section class="xp-info">
          <h2 class="section-title">Progress</h2>
          <div class="info-card">
            <p><strong>Level:</strong> 5</p>
            <p><strong>Title:</strong> Dishwashing Master</p>
            <div class="progress-bar-container">
              <div class="progress-bar" :style="{ width: xpProgress + '%' }"></div>
            </div>
            <p class="xp-text">XP: 450 / 500</p>
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
      </main>
    </div>
  </template>
  
  <script>
  export default {
    name: "UserProfilePage",
    data() {
      return {
        xpProgress: 90, // Percentage of XP towards the next level
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
  </style>
  