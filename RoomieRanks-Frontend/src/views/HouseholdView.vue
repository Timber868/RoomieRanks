<template>
    <div id="household-page">
      <!-- Household Page for Users Without a Household -->
      <main class="household-container">
        <h1 class="page-title">Manage Your Household</h1>
        <div class="button-container">
          <button class="household-button create-button">Create Household</button>
          <button class="household-button join-button">Join Existing Household</button>
        </div>
      </main>
    </div>
  </template>
  
  <script>
  export default {
    name: "HouseholdPageWithoutHousehold",
  };
  </script>
  
  <style scoped>
  /* General Reset */
  #household-page {
    margin: 0;
    padding: 2rem 0;
    font-family: 'Roboto', sans-serif;
    color: #fff;
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  
  /* Page Container */
  .household-container {
    text-align: center;
    max-width: 600px;
    width: 90%;
    background: rgba(255, 255, 255, 0.05);
    padding: 2rem;
    border-radius: 15px;
    box-shadow: 0 6px 15px rgba(0, 0, 0, 0.3);
  }
  
  .page-title {
    font-size: 2.5rem;
    margin-bottom: 2rem;
    color: #fca400;
  }
  
  .button-container {
    display: flex;
    gap: 1.5rem;
    justify-content: center;
  }
  
  .household-button {
    padding: 1rem 2rem;
    font-size: 1.2rem;
    font-weight: bold;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    transition: transform 0.2s ease, box-shadow 0.2s ease;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
  }
  
  .household-button:hover {
    transform: translateY(-3px);
    box-shadow: 0 6px 15px rgba(0, 0, 0, 0.4);
  }
  
  .create-button {
    background: linear-gradient(90deg, #fca400, #ffcd3c);
    color: #000;
  }
  
  .join-button {
    background: linear-gradient(90deg, #1e3a8a, #2563eb);
    color: #fff;
  }
  </style>
  