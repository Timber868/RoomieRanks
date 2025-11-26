<template>
  <div class="login">
      <p class="title">Login</p>
      <form class="form">
          <div class="input">
              <label for="username">Username</label>
              <input type="text" name="username" id="username" placeholder="" required v-model="username">
          </div>
          <div class="input">
              <label for="password">Password</label>
              <input type="password" name="password" id="password" placeholder="" required v-model="password">
          </div>
          <button class="sign-in" @click="attemptLogin" v-bind:disabled="!isInputValid()">Sign in</button>
      </form>
      <p class="create-account">Don't have an account?
          <RouterLink to="/Register" class="nav-item">Sign up</RouterLink>
      </p>
  </div>
</template>

<script>
import axios from "axios";
import { session } from "../session.ts";
import { formatErrorMessage } from '../utils/errorFormatter.ts';

const axiosClient = axios.create({
  //NOTE: it's baseURL, not baseUrl
  baseURL: "http://localhost:8080/api/v1"
});

export default {
  name: 'LoginView',

  data() {
  return {
    username: null,
    password: null,
    permissionLevel: 0
  };
},
  methods: {
  async attemptLogin() {
          event.preventDefault();  // Prevent the form from submitting and updating the URL

    const newLoginDTO = {
      username: this.username,
      password: this.password,
    };
    try {
              console.log(sessionStorage.getItem("loggedInUsername"))
              console.log(newLoginDTO)
              console.log(this.username)
              const response = await axiosClient.post("/login", newLoginDTO, {
                  params: { loggedInUsername: sessionStorage.getItem("loggedInUsername") }   // Add the query parameter
              });
              // The backend only returns a success message, not user data
              // Store the username before clearing inputs
              const loggedInUsername = this.username;
              sessionStorage.setItem("loggedInUsername", loggedInUsername);
              sessionStorage.setItem("permissionLevel", 1);
              this.clearInputs();
              session.updateSession(loggedInUsername, 1); // Update global state
              console.log("loggedInUsername is now:", sessionStorage.getItem("loggedInUsername"));
              console.log("permissionLevel is now:", sessionStorage.getItem("permissionLevel"));
              
              // Show success popup
              this.$swal({
                title: 'Success!',
                text: 'You have been logged in successfully.',
                icon: 'success',
                confirmButtonText: 'OK',
                confirmButtonColor: '#fc8400'
              }).then(() => {
                this.$router.push("/");
              });
    }
    catch (error) {
              console.log("hey")
              // Check if the error is a server response with a status code
              if (error.response) {
                  const status = error.response.status;
                  // Backend returns errors in format: { "error": "error message" }
                  const message = error.response.data?.error || "An error occurred.";
                  
                                     // Format the error message: capitalize first letter and add period if missing
                   let errorMessage = formatErrorMessage(message);
                  console.log("Backend error:", message);
                  
                  // Show error popup
                  this.$swal({
                    title: 'Login Error!',
                    text: errorMessage,
                    icon: 'error',
                    confirmButtonText: 'OK',
                    confirmButtonColor: '#dc2626'
                  });
              } else {
                  // Network or unexpected error
                  console.error(error);
                  this.$swal({
                    title: 'Connection Error!',
                    text: 'Unable to connect to the server. Please check your internet connection and try again.',
                    icon: 'error',
                    confirmButtonText: 'OK',
                    confirmButtonColor: '#dc2626'
                  });
              }
          }


  },
  clearInputs() {
    this.username = null,
    this.password = null,
          this.permissionLevel = 0
  },
  isInputValid() {
    return this.username   
              && this.password
  }

  }
}
</script>

<!--Open source: https://uiverse.io/Yaya12085/short-panda-24-->
<style scoped>
.login {
width: 500px;
margin-top: 100px;
margin-left: auto;
margin-right: auto;
border-radius: 0.75rem;
background-color: #1e1e1e;
padding: 2rem;
color: rgba(243, 244, 246, 1);
}

.title {
text-align: center;
font-size: 1.5rem;
line-height: 2rem;
font-weight: 700;
}

.form {
margin-top: 1.5rem;
}

.input {
margin-top: 0.25rem;
font-size: 0.875rem;
line-height: 1.25rem;
}

.input label {
display: block;
color: rgba(156, 163, 175, 1);
margin-bottom: 4px;
margin-top: 20px;
}

.input input {
width: 100%;
border-radius: 0.375rem;
border: 1px solid rgba(55, 65, 81, 1);
outline: 0;
background-color: #1e1e1e;
padding: 0.75rem 1rem;
color: rgba(243, 244, 246, 1);
}

.input input:focus {
border-color: rgba(252, 164, 0, 0.5);
}

.signup a {
color: rgba(243, 244, 246, 1);
text-decoration: none;
font-size: 14px;
}

.signup a:hover {
text-decoration: underline rgba(252, 164, 0, 0.5);
}

.sign-in:disabled {
background-color: rgba(75, 85, 99, 1);
color: rgba(243, 244, 246, 0.5);
cursor: not-allowed;
}

.sign-in {
margin-top: 10px;
display: block;
width: 100%;
background-color: rgba(252, 164, 0, 0.5);
padding: 0.75rem;
text-align: center;
color: rgba(17, 24, 39, 1);
border: none;
border-radius: 0.375rem;
font-weight: 600;
cursor: pointer;
}

.sign-in:hover:not(:disabled) {
background-color: rgba(252, 164, 0, 0.7);
}

.create-account {
margin-top: 10px;
text-align: center;
font-size: 0.75rem;
line-height: 1rem;
color: rgba(156, 163, 175, 1);
}
</style>
  