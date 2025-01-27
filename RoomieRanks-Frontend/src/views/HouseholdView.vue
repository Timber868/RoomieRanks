<template>
  <div id="household-page">
    <!-- Household Page for Users Without a Household -->
    <main class="household-container">
      <section v-if="buttonWasPressed != true">
        <h1 class="page-title">Manage Your Household</h1>
        <div class="button-container">
          <button class="household-button create-button" @click="buttonWasPressed = true">Create Household</button>
          <button class="household-button join-button">Join Existing Household</button>
        </div>
      </section>
      <section v-if="buttonWasPressed == true">
        <div class="button-container">
          <form @submit.prevent="createHousehold" class="household-form">

          <div class="form-group">
            <label for="householdText">Household Name:</label>
            <textarea id="householdText" v-model="householdText" required></textarea>
          </div>

          <div class="form-buttons">
            <button type="submit" class="submit-household-button">Create Household</button>
            <button type="button" class="cancel-household-button" @click="handleCancel">Cancel</button>
          </div>
        </form>
        </div>
      </section>
    </main>
  </div>
</template>
  
<script>
import axios from "axios";
import { session } from "../session.ts";

const axiosClient = axios.create({
  //NOTE: it's baseURL, not baseUrl
  baseURL: "http://localhost:8080/api/v1"
});

export default {
  name: "HouseholdPageWithoutHousehold",

  data() {
    return {
      buttonWasPressed: false,
      username: session.loggedInUsername,

      // Household creation form data
      householdText: "",
    };
  },
  methods: {
  async createHousehold() {
    console.log("Creating household with name:", this.householdText);
    if (this.householdText.length === 0) {
      alert("Please enter a household name.");
      return;
    }

    const householdRequest = {
      name: this.householdText,
    };

    try {
      // 1) Wait for the POST request to finish
      const response = await axiosClient.post("/household", householdRequest);
      // The server’s data will be in response.data
      console.log("POST /household response:", response);

      // 2) Extract the new household's ID from response.data
      //    (Adjust this if your API uses a different field name)
      const houseHoldID = response.data;
      console.log("Household ID:", houseHoldID);
      console.log("Username:", sessionStorage.getItem("loggedInUsername"));

      // 3) Update the user’s household ID with PUT
      const response2 = await axiosClient.put(
        "/user/" + sessionStorage.getItem("loggedInUsername") + "/household/" + houseHoldID
      );
      console.log("PUT /user/... response:", response2);

      // 4) Reset form and hide it
      this.householdText = "";
      this.buttonWasPressed = false;
    } catch (error) {
      console.error("Error creating household:", error);
    }
  },

  handleCancel() {
    // Reset the form
    this.householdText = "";
    // Hide the form
    this.buttonWasPressed = false;
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

  /* Overall page container */
#household-page {
  background-color: #121212; /* Dark background */
  color: #ffffff;           /* White text */
  padding: 2rem;
  min-height: 100vh;
}

/* Main container styling */
.household-container {
  max-width: 800px;
  margin: 0 auto;
  background-color: #1e1e1e;  /* Slightly lighter than the page background */
  border-radius: 8px;
  padding: 2rem;
}

/* Title at the top */
.page-title {
  font-size: 1.8rem;
  margin-bottom: 1.5rem;
  text-align: center;
}

/* Buttons container */
.button-container {
  display: flex;
  gap: 1rem;
  justify-content: center;
  margin-bottom: 2rem;
}

/* General button styling */
.household-button {
  border: none;
  border-radius: 6px;
  padding: 0.75rem 1.5rem;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
  color: #ffffff;
  background-color: #3a3a3a; /* Default dark gray */
  transition: background-color 0.2s ease;
}

.household-button:hover {
  background-color: #555555;
}

/* Optionally differentiate the create & join buttons */
.create-button {
  background-color: #4C3F2E; /* A warm brown tone, similar to "Your Chores" cards */
}

.create-button:hover {
  background-color: #5a4b38;
}

.join-button {
  background-color: #003947; /* A deep teal tone, similar to "Housemates' Chores" cards */
}

.join-button:hover {
  background-color: #004b5c;
}

/* Form container & elements */
.household-form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  background-color: #1e1e1e; /* Matches the household-container, or make it a shade lighter if you prefer */
  padding: 1.5rem;
  border-radius: 8px;
}

/* Label + textarea grouping */
.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.form-group label {
  font-weight: 600;
}

.form-group textarea {
  resize: none;
  width: 100%;
  height: 3rem;
  border-radius: 4px;
  border: none;
  padding: 0.5rem;
  font-size: 1rem;
  color: #ffffff;
  background-color: #2a2a2a;
}

/* Form buttons */
.form-buttons {
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
}

.submit-household-button,
.cancel-household-button {
  border: none;
  border-radius: 4px;
  padding: 0.75rem 1.25rem;
  cursor: pointer;
  font-size: 1rem;
  color: #ffffff;
  background-color: #3a3a3a;
  transition: background-color 0.2s ease;
}

.submit-household-button:hover,
.cancel-household-button:hover {
  background-color: #555555;
}

.submit-household-button {
  background-color: #4C3F2E; /* Same as .create-button for consistency */
}

.submit-household-button:hover {
  background-color: #5a4b38;
}

.cancel-household-button {
  background-color: #444444; /* A neutral gray for "Cancel" */
}

  </style>
  