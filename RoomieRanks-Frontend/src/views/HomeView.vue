<template>
  <div id="home">
    <!-- Chore List Section -->
    <main class="chores-container">
      <!-- User's Chores -->
      <section class="chores-section">
        <h2 class="section-title">Your Chores</h2>
        <div v-for="chore in userChores" :key="chore.id" class="chore-card user-chore">
          <div class="chore-details">
            <h3 class="chore-name">{{ chore.name }}</h3>
            <p class="chore-due-date">Due: {{ chore.dueDate }}</p>
            <p class="chore-status">Difficulty: {{ chore.difficulty }}</p>
            <p class="chore-status">Time Estimate: {{ chore.timeEstimate }}</p>
            <p class="chore-status">Status: {{ chore.status }}</p>
          </div>
          <div class="chore-assignee">Assigned to: {{ chore.assignee }}</div>
        </div>
      </section>

      <!-- Housemates' Chores -->
      <section class="chores-section">
        <h2 class="section-title">Housemates' Chores</h2>
        <div v-for="chore in housemateChores" :key="chore.id" class="chore-card housemate-chore">
          <div class="chore-details">
            <h3 class="chore-name">{{ chore.name }}</h3>
            <p class="chore-due-date">Due: {{ chore.dueDate }}</p>
            <p class="chore-status">Difficulty: {{ chore.difficulty }}</p>
            <p class="chore-status">Time Estimate: {{ chore.timeEstimate }}</p>
            <p class="chore-status">Status: {{ chore.status }}</p>
          </div>
          <div class="chore-assignee">Assigned to: {{ chore.assignee }}</div>
        </div>
      </section>
    </main>
  </div>
</template>

<script>
export default {
  name: "HomePage",
  data() {
    return {
      userChores: [
        { id: 1, name: "Take out trash", dueDate: "2025-01-28", status: "On Going", assignee: "You", difficulty: "Easy", timeEstimate: "5 minutes" },
        { id: 2, name: "Wash dishes", dueDate: "2025-01-29", status: "Completed", assignee: "You", difficulty: "Medium", timeEstimate: "15 minutes" },
      ],
      housemateChores: [
        { id: 3, name: "Vacuum living room", dueDate: "2025-01-30", status: "On Going", assignee: "Alex", difficulty: "Hard", timeEstimate: "30 minutes" },
        { id: 4, name: "Mow the lawn", dueDate: "2025-02-01", status: "On Going", assignee: "Jamie", difficulty: "Hard", timeEstimate: "1 hour" },
      ],
    };
  },
};
</script>

<style scoped>
/* General Reset */
#home {
  margin: 0;
  padding: 0;
  font-family: 'Roboto', sans-serif;
}

/* Chores Container */
.chores-container {
  max-width: 1200px;
  margin: 5rem auto;
  padding: 0 1rem;
}

.chores-section {
  margin-bottom: 2rem;
}

.section-title {
  font-size: 1.5rem;
  margin-bottom: 1.5rem;
  color: #fff;
}

/* Chore Cards */
.chore-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #ffffff;
  border-radius: 8px;
  padding: 1rem;
  margin-bottom: 1rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s ease;
}

.chore-card:hover {
  transform: scale(1.02);
}

.user-chore {
  border-left: 5px solid #fca400;
  background-color: rgba(252, 164, 0, 0.1);
}

.housemate-chore {
  border-left: 5px solid #00adef;
  background-color: rgba(0, 173, 239, 0.1);
}

.chore-details {
  flex: 1;
}

.chore-name {
  font-size: 1.25rem;
  margin: 0 0 0.5rem;
  color: #fff;
}

.chore-due-date,
.chore-status {
  margin: 0;
  font-size: 0.95rem;
  color: #fff;
}

.chore-assignee {
  font-weight: bold;
  font-size: 1rem;
  margin-left: 1rem;
  color: #fff;
}
</style>

