
<body>
    <p>
            This documentantion was writen for <a href="[https://github.com/maxogod](https://github.com/maxogod/Front-End-Projects)" target="_blank">this little project</a>
    </p>
    <nav id="navbar">
        <header id="navbar-header">AlgoGram Documentation</header>
        <ul>
            <li>
                <a href="#Introduction" class="nav-link">Introduction</a>
            </li>
            <li>
                <a href="#How_to_use" class="nav-link">How to use</a>
            </li>
            <li>
                <a href="#Commands" class="nav-link">Commands</a>
            </li>
            <li>
                <a href="#Used_data_structures" class="nav-link">Used data structures</a>
            </li>
            <li>
                <a href="#New_data_types" class="nav-link">New data types</a>
            </li>
            <li>
                <a href="#Reference" class="nav-link">Reference</a>
            </li>
        </ul>
    </nav>
    <main id="main-doc">
      <section class="main-section" id="Introduction">
        <header>~Introduction</header>
        <p>
          This documentation page is for FreeCodeCamp's responsive web design
          curriculum; It is in English even though the mentioned repository is
          in Spanish.
        </p>
        <p>
          AlgoGram is basically a very basic infrastructure of a social media
          platform, it's real purpose was to practice usage of several
          data-structures listed below.
        </p>
      </section>
      <section class="main-section" id="How_to_use">
        <header>~How to use</header>
        <p>
          In order to run this aplication you will need to execute the following
          command in the terminal:
        </p>
        <code>
          ./algogram &lt;txt file containing the users of the app&gt;
        </code>
      </section>
      <section class="main-section" id="Commands">
        <header>~Commands</header>
        <p>All the executable commands during stdin are:</p>
        <p>login Logs a user in (wow)</p>
        <code> login &lt;user-name&gt; </code>
        <p>logout Logs the user that's currently logged in</p>
        <code> logout </code>
        <p>publicar Posts a new message</p>
        <code> publicar &lt;post-text&gt; </code>
        <p>ver_siguiente_feed Shows the user the next post in the feed</p>
        <code> ver_siguiente_feed </code>
        <p>likear_post Likes a post based on it's id</p>
        <code> likear_post &lt;post-id&gt; </code>
        <p>mostrar_likes Shows the list of people who liked a certain post</p>
        <code> mostrar_likes &lt;post-id&gt; </code>
      </section>
      <section class="main-section" id="Used_data_structures">
        <header>~Used data structures</header>
        <p>Data structures that were used in this project are:
            <ul>
                <li>Priority queue for the feed</li>
                <li>Hashmaps for the posts in order to get them by id's and same goes for the users</li>
                <li>Binary search tree to save the list of people who liked a post in alfabetic order</li>
            </ul>
        </p>
      </section>
      <section class="main-section" id="New_data_types">
        <header>~New data types</header>
        <p>New data types that were implemented are just two:
            <ul>
                <li>Users data type</li>
                <li>Posts data type</li>
            </ul>
            both have a a fitting behavior
        </p>
      </section>
      <section class="main-section" id="Reference">
        <header>~Reference</header>
        <p>
            The repository is currently private but will be public soon in <a href="https://github.com/maxogod" target="_blank">this github account</a>
        </p>
      </section>
    </main>
  </body>
</html>
