# COOKIT

<div align="center">
  <a href="">
    <img src="" width="304" height="297">
  </a>

  <p align="center">
    Capstone Program Immersive Alterra Academy
    <br />
    <a href="https://app.swaggerhub.com/apis-docs/STARCON10_1/ALTA-Cookit-BE/1.0"><strong>| Open API Documentation |</strong></a>
    <br />
    <br />
  </p>
</div>

## 🧑‍💻 About the Project

<p align="justify">COOKIT is a web application-based social media that is useful for finding food recipes online. You can get food recipes from all over the world at COOKIT. Users can not only upload homemade recipes but can also upload re-cooking results. Users can also like recipes and follow other users. Users can also sell ingredients at COOKIT, but only verified users can sell them.</p>

## 🛠 Technology Stack

<div align="center">
<img src="techno_stack.png">
  </div>

# 🔗 ERD

<div align="center">
<img src="cookit_ERD.png" width="800" height="800">
  </div>

# ⚡ Features

<details>
  <summary>🎫 Auth</summary>
  
| Method      | Endpoint            | Params      |q-Params            | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- |--------------------| ----------- | --------------------------------------- |
| POST        | /register           | -           |-                   | NO         | Register a new User                |
| POST        | /login      | -           |-                   | NO         | Login to the system        |
  
</details>

<details>
  <summary>🙍‍♂️ Users</summary>
  
| Method      | Endpoint            | Params      |q-Params            | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- |--------------------| ----------- | --------------------------------------- |
| GET        | /users           | -           |-                   | YES         | Show profile                |
| PUT        | /users      | -           |-                   | YES         | Update profile data        |
| DELETE        | /users      | -           |-                   | YES         | Delete user data        |
| GET        | /users/search      | -           |-                   | YES         | Search another users with username        |
| PUT        | /users/password      | -           |-                   | YES         | Update password account        |
| GET        | /users/(ID)      | -           |- ID                  | YES         | Show another user profile        |
| GET        | /users/follower      | -           |-                   | YES         | Show list follower        |
| GET        | /users/following      | -           |-                   | YES         | Show list following        |
| GET        | /users/upgrade      | -           |-                   | YES         | Request upgrade account        |
</details>

<details> 
    <summary>👮 Admin </summary>
| Method      | Endpoint            | Params      |q-Params            | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- |--------------------| ----------- | --------------------------------------- |
| GET        | /users/listverify           | -           |-                   | YES         | Show list for user request upgrading account                |
| PUT        | /users/approval/(ID)      | -           |- ID                  | YES         | Accepting or deny user request upgrade account for admin        |
</details>

<details> 
    <summary>🙋‍♂️ Followers </summary>
| Method      | Endpoint            | Params      |q-Params            | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- |--------------------| ----------- | --------------------------------------- |
| POST        | /users/follow/(ID)           | -           |- ID                  | YES         | Following another user                |
| DELETE        | /users/unfollow/(ID)      | -           |- ID                  | YES         | Unfollow users        |
</details>

<details> 
    <summary>🍳 Recipes</summary>
| Method      | Endpoint            | Params      |q-Params            | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- |--------------------| ----------- | --------------------------------------- |
| GET        | /recipes           | -           |-                   | YES         | Show list recepies                |
| POST        | /recipes      | -           |-                   | YES         | Insert new recipe        |
| PUT        | /recipes(ID)      | -           |-  ID                 | YES         | Update recipe        |
| DELETE        | /recipes(ID)      | -           |-  ID                 | YES         | Delete recipe by ID        |
| GET        | /users/recipes/timeline      | -           |-                   | YES         | Show timeline recipes        |
| GET        | /recipes/trending      | -           |-                   | YES         | Show trending recipes        |
| GET        | /recipes/(ID)/detail      | -           |-  ID                 | YES         | Show detail recipes        |
| POST        | /recipes/(ID)/like      | -           |-  ID                 | YES         | Like recipes        |
| DELETE        | /recipes/(ID)/unlike      | -           |-  ID                 | YES         | Unlike recipes        |
</details>

<details> 
    <summary>🖼️ Images</summary>
| Method      | Endpoint            | Params      |q-Params            | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- |--------------------| ----------- | --------------------------------------- |
| POST        | /recipes/(recipe_id)/images           | -           |- recipe_id                  | YES         | Insert new recipes image                |
| DELETE        | /recipes/(recipe_id)/images      | -           |- recipe_id                  | YES         | Delete recipes image        |
| PUT        | /recipes/(recipe_id)/images/(image_id)      | -           |- recipe_id and image_id                 | YES         | Update recipes image        |
| DELETE        | /recipes/(recipe_id)/images/(image_id)      | -           |- recipe_id and image_id                  | YES         | Delete recipes image        |
</details>