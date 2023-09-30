Introduction

Htmx is a JavaScript library for performing AJAX requests, triggering CSS transitions, and invoking WebSocket and server-sent events directly from HTML elements. Htmx lets you build modern and powerful user interfaces with simple markups.



This library weighs ~10KB (min.gz’d), it is dependency-free (i.e., it does not require any other JavaScript package to run), and it’s also compatible with IE11.



In this tutorial, we will be exploring the powerful features of htmx while covering the following sections:

    Installing htmx
    Sending AJAX requests with htmx
    Custom htmx input validation
    Triggering CSS animation with htmx


Installing htmx

You can get started with htmx by downloading the htmx source file or including its CDN directly in your markup, like below:

<script src="https://unpkg.com/htmx.org@1.3.3"></script>

The script above loads the current stable version of htmx, which as of writing this is version 1.3.3, on your webpage. Once that’s done, you can implement htmx features on your webpage.


Sending AJAX requests with htmx

Htmx provides a set of attributes that allows you to send AJAX requests directly from an HTML element. Available attributes include:

    hx-get — send GET request to the provided URL
    hx-post — send POST request to the provided URL
    hx-put — send PUT request to the provided URL
    hx-patch — send PATCH request to the provided URL
    hx-delete — send DELETE request to the provided URL

Code sample

<button hx-get="http://localhost/todos">Load Todos</button>

The code example above tells the browser that when the user clicks the button, it sends a GET request (hx-get) to the provided URL, which in this case is http://localhost/todos.



Elijah Asaolu
I am a programmer, I have a life.
Htmx: The newest old way to make web apps

June 4, 2021 6 min read
HTMX: The Newest Old Way to Make Web Apps
See how LogRocket's AI-powered error tracking works
no signup required
Check it out
Introduction

Htmx is a JavaScript library for performing AJAX requests, triggering CSS transitions, and invoking WebSocket and server-sent events directly from HTML elements. Htmx lets you build modern and powerful user interfaces with simple markups.

This library weighs ~10KB (min.gz’d), it is dependency-free (i.e., it does not require any other JavaScript package to run), and it’s also compatible with IE11.

In this tutorial, we will be exploring the powerful features of htmx while covering the following sections:

    Installing htmx
    Sending AJAX requests with htmx
    Custom htmx input validation
    Triggering CSS animation with htmx

Installing htmx

You can get started with htmx by downloading the htmx source file or including its CDN directly in your markup, like below:

<script src="https://unpkg.com/htmx.org@1.3.3"></script>

The script above loads the current stable version of htmx, which as of writing this is version 1.3.3, on your webpage. Once that’s done, you can implement htmx features on your webpage.
Sending AJAX requests with htmx

Htmx provides a set of attributes that allows you to send AJAX requests directly from an HTML element. Available attributes include:

    hx-get — send GET request to the provided URL
    hx-post — send POST request to the provided URL
    hx-put — send PUT request to the provided URL
    hx-patch — send PATCH request to the provided URL
    hx-delete — send DELETE request to the provided URL

Code sample

<button hx-get="http://localhost/todos">Load Todos</button>

The code example above tells the browser that when the user clicks the button, it sends a GET request (hx-get) to the provided URL, which in this case is http://localhost/todos.

Htmx get-request

By default, the response returned from any htmx request will be loaded in the current element that is sending the request. In the In the Targeting elements for AJAX requests section, we will be exploring how to load the response in another HTML element.

Targeting elements for AJAX requests section, we will be exploring how to load the response in another HTML element.


Trigger requests

You should note that AJAX requests in htmx are triggered by the natural event of the element. For example, input, select, and textarea are triggered by the onchange event, and form is triggered by the onsubmit event, and every other thing is triggered by the onclick event.

In a situation where you want to modify the event that triggers the request, htmx provides a special hx-trigger attribute for this:

<div hx-get="http://localhost/todos" hx-trigger="mouseenter">
    Mouse over me!
</div>

In the example above, the GET request will be sent to the provided URL if and only if the user’s mouse hovers on the div.
Trigger modifiers

The hx-trigger attribute mentioned in the previous section accepts an additional modifier to change the behavior of the trigger. Available trigger modifiers include:

    once — ensures a request will only happen once
    changed — issues a request if the value of the HTML element has changed
    delay:<time interval> — waits for the given amount of time before issuing the request (e.g., delay-1s). If the event triggers again, the countdown is reset
    throttle:<time interval> — waits the given amount of time before sending the request (e.g., throttle:1s). But unlike delay, if a new event occurs before the time limit is reached, the event will be in a queue so that it will trigger at the end of the previous event
    from:<CSS Selector> — listens for the event on a different element

Code sample

<input
    type="text"
    hx-get="http://localhost/search"
    hx-trigger="keyup changed delay:500ms" />

In the code sample provided above, once the user performs a keyup event on the input element (i.e., the user types any text in the input box) and its previous value changes, the browser will automatically send a GET request to http://localhost/search after 500ms.
Polling with the htmx-trigger attribute

In the htmx-trigger attribute, you can also specify every n seconds rather than waiting for an event that triggers the request. With this option, you can send a request to a particular URL every n seconds:

  <div hx-get="/history" hx-trigger="every 2s">
  </div>

The code sample above tells the browser to issue a GET request to /history endpoint every 2s and load the response into the div.
Targeting elements for AJAX requests

In previous sections, we’d mentioned that the response from an AJAX request in htmx will be loaded into the element making the request. If you need the response to be loaded into a different element, you can use the hx-target attribute to do this. This attribute accepts a CSS selector and automatically injects the AJAX response into an HTML element with the specified selector.