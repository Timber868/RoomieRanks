<template>
  <div class="create-account">
  <p class="title">Sign Up</p>
  <form class="form">
      <div class="input">
          <label for="name">Name</label>
          <input type="text" name="name" id="name" placeholder="" required v-model="name">
      </div>
      <div class="input">
          <label for="username">Username</label>
          <input type="text" name="username" id="username" placeholder="" required v-model="username">
      </div>
      <div class="input">
          <label for="password">Password</label>
          <input type="password" name="password" id="password" placeholder="" required v-model="password">
      </div>
      <div class="input">
          <label for="email">Email</label>
          <input type="text" name="email" id="email" placeholder="" required v-model="email">
      </div>
          <button class="sign-in nav-item" @click="attemptSignUp" v-bind:disabled="!isInputValid()">Create Account</button>
          <p class="login">Already have an account?
          <RouterLink to="/Login" class="nav-item">Login</RouterLink>
          </p>
  </form>
  </div>
</template>

<script>
import axios from "axios";
import { session } from '../session.ts'
import { formatErrorMessage } from '../utils/errorFormatter.ts'


const axiosClient = axios.create({
// NOTE: it's baseURL, not baseUrl
baseURL: "http://localhost:8080/api/v1"
});

export default {
  name: 'SignUpView',

  data() {
  return {
          name: null,
    username: null,
    password: null,
          email: null
  };
},

  methods: {
      async attemptSignUp() {
          event.preventDefault();  // Prevent the form from submitting and updating the URL

    const newUserRequest = {
      username: this.username,
              email: this.email,
              password: this.password,
              name: this.name,
    };
    try {
              console.log(sessionStorage.getItem("loggedInUsername"));
              console.log(newUserRequest)
      const response = await axiosClient.post("/register", newUserRequest);
              this.username = response.data.username;
              this.permissionLevel = response.data.permissionLevel;
              sessionStorage.setItem("loggedInUsername", this.username);
              sessionStorage.setItem("permissionLevel", 1);
              this.clearInputs();
              session.updateSession(response.data.username, response.data.permissionLevel);
              console.log("loggedInUsername is now:", sessionStorage.getItem("loggedInUsername"));
              console.log("permissionLevel is now:", sessionStorage.getItem("permissionLevel"));
              
              // Show success popup
              this.$swal({
                title: 'Success!',
                text: 'Your account has been created successfully.',
                icon: 'success',
                confirmButtonText: 'OK',
                confirmButtonColor: '#fc8400'
              }).then(() => {
                this.$router.push("/profile");
              });
    }
    catch (error) {
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
                    title: 'Error!',
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
          && this.name
          && this.email
      }
  }
}

</script>

<style scoped>
.create-account {
width: 500px;
margin-top: 100px;
margin-left: auto;
margin-right: auto;
border-radius: 10px;
background-color: #1e1e1e;
padding: 2rem;
color: rgba(243, 244, 246, 1);
}

.title {
text-align: center;
font-size: 1.5rem;
line-height: 2rem;
font-weight: 700;
color: rgba(243, 244, 246, 1);
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
border-color:rgba(252, 164, 0, 0.5);
}

.signup a {
color: rgba(243, 244, 246, 1);
text-decoration: none;
font-size: 14px;
}

.signup a:hover {
text-decoration: underline rgba(252, 164, 0, 0.5);
}

.sign-in {
margin-top: 20px;
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

.sign-in:hover {
text-decoration: underline rgba(55, 65, 81, 1);
}

.sign-in:disabled {
background-color: rgba(75, 85, 99, 1);
color: rgba(243, 244, 246, 0.5);
cursor: not-allowed; 
}

.login {
margin-top: 10px;
text-align: center;
font-size: 0.75rem;
line-height: 1rem;
color: rgba(156, 163, 175, 1);
}
</style>
  