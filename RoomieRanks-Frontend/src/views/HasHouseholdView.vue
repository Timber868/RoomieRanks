<template>
  <div id="household-page">
    <!-- Household Page for Users with a Household -->
    <main class="household-container">
      <!-- Household Info Card -->
      <div class="household-info-card">
        <h2 class="household-name">The Clean Team</h2>
        <p class="household-id">ID: 12345</p>
      </div>

      <!-- Chore List Section -->
      <section class="chores-section">
        <h2 class="section-title">Chores</h2>
        <ul class="chore-list">
          <li v-for="chore in chores" :key="chore.id" class="chore-card">
            <div class="chore-details">
              <p class="chore-name"><strong>Name:</strong> {{ chore.name }}</p>
              <p class="chore-difficulty"><strong>Difficulty:</strong> {{ chore.difficulty }}</p>
              <p class="chore-time"><strong>Time Estimate:</strong> {{ chore.timeEstimate }} mins</p>
            </div>
          </li>
        </ul>
        <button class="add-chore-button" @click="toggleAddChore">Add Chore</button>
      </section>

      <!-- Add Chore Form -->
      <section v-if="showAddChore" class="add-chore-form">
        <h3 class="form-title">Add a New Chore</h3>
        <form @submit.prevent="addChore">
          <div class="form-group">
            <label for="chore-name">Name</label>
            <input type="text" id="chore-name" v-model="newChore.name" required />
          </div>
          <div class="form-group">
            <label for="chore-difficulty">Difficulty</label>
            <select id="chore-difficulty" v-model="newChore.difficulty" required>
              <option value="Easy">Easy</option>
              <option value="Medium">Medium</option>
              <option value="Hard">Hard</option>
            </select>
          </div>
          <div class="form-group">
            <label for="chore-time">Time Estimate (mins)</label>
            <input type="number" id="chore-time" v-model="newChore.timeEstimate" required />
          </div>
          <button type="submit" class="submit-button">Submit</button>
        </form>
      </section>
    </main>
  </div>
</template>

<script>
export default {
  name: "HouseholdPageWithHousehold",
  data() {
    return {
      chores: [
        { id: 1, name: "Vacuum Living Room", difficulty: "Medium", timeEstimate: 30 },
        { id: 2, name: "Wash Dishes", difficulty: "Easy", timeEstimate: 15 },
        { id: 3, name: "Mow Lawn", difficulty: "Hard", timeEstimate: 60 },
      ],
      showAddChore: false,
      newChore: {
        name: "",
        difficulty: "Easy",
        timeEstimate: "",
      },
    };
  },
  methods: {
    toggleAddChore() {
      this.showAddChore = !this.showAddChore;
    },
    addChore() {
      const newId = this.chores.length + 1;
      this.chores.push({ id: newId, ...this.newChore });
      this.newChore = { name: "", difficulty: "Easy", timeEstimate: "" };
      this.showAddChore = false;
    },
  },
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
  padding-bottom: 4rem;
}

/* Household Info Card */
.household-info-card {
  text-align: center;
  background: linear-gradient(90deg, #fca400, #ffcd3c);
  color: #000;
  border-radius: 8px;
  padding: 2rem;
  margin: 3rem auto 2rem;
  max-width: 600px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
}

.household-name {
  font-size: 1.8rem;
  font-weight: bold;
  margin: 0;
}

.household-id {
  font-size: 1rem;
  margin-top: 0.5rem;
}

/* Chores Section */
.chores-section {
  max-width: 800px;
  margin: 2rem auto 0;
  text-align: center;
}

.section-title {
  font-size: 1.8rem;
  margin-bottom: 1.5rem;
  color: #fca400;
}

.chore-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: rgba(252, 164, 0, 0.1);
  border-left: 5px solid #fca400;
  border-radius: 8px;
  padding: 0.8rem;
  margin-bottom: 1rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s ease;
}

.chore-card:hover {
  transform: scale(1.02);
}

.chore-details {
  flex: 1;
}

.chore-name {
  font-size: 1.2rem;
  margin: 0 0 0.5rem;
  color: #fff;
}

.chore-difficulty,
.chore-time {
  margin: 0;
  font-size: 1rem;
  color: #fff;
}

/* Add Chore Button */
.add-chore-button {
  padding: 1rem 2rem;
  font-size: 1.2rem;
  font-weight: bold;
  background: linear-gradient(90deg, #fca400, #ffcd3c);
  color: #000;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
}

.add-chore-button:hover {
  transform: translateY(-3px);
  box-shadow: 0 6px 15px rgba(0, 0, 0, 0.4);
}

/* Add Chore Form */
.add-chore-form {
  max-width: 600px;
  margin: 0 auto;
  background: rgba(255, 255, 255, 0.1);
  padding: 1.5rem;
  border-radius: 10px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
}

.form-title {
  font-size: 1.5rem;
  margin-bottom: 1rem;
  color: #fca400;
}

.form-group {
  margin-bottom: 1rem;
  text-align: left;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: #e5e7eb;
}

.form-group input,
.form-group select {
  width: 100%;
  padding: 0.8rem;
  border: none;
  border-radius: 5px;
  font-size: 1rem;
  margin-bottom: 0.5rem;
}

.submit-button {
  padding: 1rem 2rem;
  font-size: 1.2rem;
  font-weight: bold;
  background: linear-gradient(90deg, #1e3a8a, #2563eb);
  color: #fff;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
}

.submit-button:hover {
  transform: translateY(-3px);
  box-shadow: 0 6px 15px rgba(0, 0, 0, 0.4);
}
</style>