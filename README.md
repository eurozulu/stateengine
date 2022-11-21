# State Engine
An experimental play with a simple state engine.  

## Concept  
Based on a common example of a vending machine, a simple state engine has three states:  

- Empty
- Locked
- Unlocked

To tranistion through a state, an Action is applied.  
There are three Actions:  

- Push Button
- Insert Coin
- Add Bottle

The machine begins in the `Empty` state.  
In Empty only the `Add Bottle` action will change the state to `Locked`.
`Push button` and `Insert coin` have no effect, returning to Empty state.

In `Locked` state only the `Insert coin` action will move to `Unlocked`,
`Push button` has no effect.  `Add Bottle` will increase the bottles but leave in
the `Locked` state.

`Insert Coin` in the `Locked` state moves to the `Unlocked` state, where `Button Push`
will move it back to Locked.
If there are still unused coins, the machine is held in `Unlocked`
state until all the coins are exhausted, before returning to `Locked` state.

`Insert coin` is rejected when the machine has run out of bottle.

### Design
The design is centered around the State.
State is a simple int/consts.

Each State has an actionMap of its own, defining which
actions it accepts and which state those actions will move to.
On an action received, the resulting mapped state becomes the next state.

Vending machine is built ontop of the State.  This contains
the state of the Vending machine:
Its coins, bottle and current state.  
On each transition of state, the vending machine checks its own status
and may modify the transition of state accordingly.  

e.g. When Push button action is given in Unlocked State,
The usual transition is to return to Locked.
If vending machine still has unused coins, it prevents this,
holding the state of unlocked until all the coins are exhuasted.

