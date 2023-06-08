# Technical Assessment: Systems Engineer

This assessment helps Vanti to hire people suited for the role of **Systems Engineer**. As part of this role you will be
developing software systems that integrate with physical hardware to provide excellent user experiences for people in
the built environment.

Working through this assessment you will develop a _device driver_ that can communicate with a physical device as part
of a wider project. Not only will this give us an idea of how you approach problems, it will also give you an idea of
the types of problems and solutions that we're trying to develop at Vanti.

There are no requirements for how you should complete this assessment, one of the things we are hoping to extract from
your submission is what a good solution looks like to you. We aren't going to penalise spelling mistakes, time taken,
how much Google is used, or how many approaches you try. Try to treat the assessment as if it were a real job, use all
the tools and resources available to you.

Once completed, we'll review the code and arrange a follow-up conversation where we'd like to discuss your approach,
andy challenges you faced, and how you overcame them - it might be useful to make notes for your own reference as you
carry out the task.

As we aren't going to be sat watching you complete this assessment, we'd like you to track your progress through the
goals via commits. You can commit as often or as little as you'd like, but we'd like to see one commit that indicates
when you feel a specific goal is completed.

## Setup and submission

This repository is read-only, you will not be able to push changes to it. We'd like you to clone the repository, and
push your changes to a code hosting site of your choosing. Try not to fork this repository directly, forks can be
tracked and we'd like to reduce easy plagiarism. Using the [GitHub importer](https://github.com/new/import) to create a
private copy is perfect!

Once your repository is available, send us a link / access, and let us know when you're done and we'll take a look.

Thank you for taking the time to complete this assessment, let's get started :rocket:

## Goals

We are attempting to develop a device driver for interrogating and controlling a Bluetooth media device, specifically
a **Q-SYS unD6IO-BT**. The device will be installed into a relaxation room at one of our client sites and allow
residents to connect their phones to the audio system installed there to play their own music via a Bluetooth
connection. The unD6IO-BT is being installed to act like a Bluetooth speaker would in a home: you activate the pairing
process on the speaker, connect via your phone, and audio comes from your phone out of the speaker.

Other members of the project team are working on the UI for the space which will take the form of a wall mounted touch
panel to provide information and control of the space: volume, radio station selection, and Bluetooth connectivity
control, etc.

### Goal 1: Allow the user to connect their phone via Bluetooth

The client does not want the unD6IO-BT to be visible on the wall of the room, it'll be located in an equipment rack
hidden from view. Because of this we need to surface functionality from the unit on the wall mounted touch panel.
Specifically we need to activate the Bluetooth pairing process on the unD6IO-BT which will be completed by the user
connecting via their phones Bluetooth controls.

Your task is to allow the UI team to trigger the unit to begin the pairing process. You do this by adding functionality
to the [`und6iobt.Driver`](./qsys/und6iobt/driver.go) type to implement the [`HasAnnounce`](./bluetooth/announce.go)
interface.

### Goal 2: Inform the user how to connect to Bluetooth

During some user testing the design team found that some people got confused with how to work the system. They want a
page on the touch panel to show help for how to connect your phone to the Bluetooth system in the room. To make this
happen they want to show the advertised Bluetooth name of the unD6IO-BT so a user knows what to look for when connecting
via their phone.

Have your driver implement the [`HasName`](./bluetooth/name.go) interface based on the friendly name of the unD6IO-BT.

### Goal 3: Show a dedicated page when a device is connected

During a design review it was noticed that we still show the _Help to connect_ page even when a phone is connected. To
fix this we need to know when a device is connected to the Bluetooth receiver and change the page shown on the UI.

You have been given the task of reporting the Bluetooth connection status via the driver. Add functionality to
the [`und6iobt.Driver`](./qsys/und6iobt/driver.go) type and implement the [`HasConnection`](./bluetooth/connection.go)
interface.

