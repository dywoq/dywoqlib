# Proposal Writing Rules

## General Rules
1. To make sure everyone understands each other, the only allowed language is **English (US/UK)**.  
2. Keep the description clear, simple, and concise. Avoid unnecessary details.  
3. Follow the required structure with the specified headers.  
4. Write in a **formal and professional** tone.  
5. Use Markdown formatting properly for readability.  
6. Each proposal should focus on **one main feature or idea** only.  
7. Avoid subjective opinions. Provide only facts, reasoning, and benefits.  
8. Provide examples when necessary.  
9. Keep consistency in naming, formatting, and code style with the rest of the project.  /
10. The naming of issues is: `[<container-name (eg. container/mapn)>] <issue-name>`
11. The implementation of features are (if it was accepted):
    - Creating your feature's branch or the fork of the repository. 
    - Writing code (implementing).
    - Writing tests and examples with the expected behavior.
    - Making pull request. The naming is: `[<container-name (eg. container/mapn)>] <issue-name>`
    - If the workflows fall, we need to fix bugs.
    - After this, pull request can be merged.
    - Your feature was accepted and implemented.

## Required Structure
Each proposal must contain the following headers:

- **Description**  
  Short explanation of the proposed feature.  

- **Features**  
  List of all new functions, structures, variables, constants, or types.  

## Features Documentation Rules
- Every feature must include both a **description** and a **signature**.  
- The format should follow standard Go inline documentation style.  

### Example
- ```algorithms.Gcd```

  Gcd finds the greatest common divisor.  

  **Signature:**  
  ```go
  func Gcd(a, b float64) float64
   ```

## Where to write
Create proposal as a sub-issue of the parent issue, if your proposal may take place in already existing package.
For example, if you're planning to create the feature inside "container" package, you may create sub-issue of https://github.com/dywoq/dywoqlib/issues/2, even it's marked as closed.
If your proposal represents new package (not sub-package), then you don't need it to create it as sub-issue.

To see the progress more comfortably, you may use the [library's Github Project](https://github.com/users/dywoq/projects/19 "").
