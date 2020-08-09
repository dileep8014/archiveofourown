import React from 'react';


const IconText = ({ icon, text }: { icon: any; text: any }) => {
  if (typeof icon === 'string') {
    return (
      <span>{icon}: {text}</span>
    );
  }

  return (
    <span>
    {React.createElement(icon, { style: { marginRight: 8 } })}
      {text}
  </span>

  );
};

export default IconText;
