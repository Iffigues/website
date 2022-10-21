function Save(obj: any, types: string, name: string) {
		const blob = new Blob([obj], { type: types });
  		const href = URL.createObjectURL(blob);
		const link = document.createElement("a");
  		link.href = href;
  		link.download =	name;
  		document.body.appendChild(link);
  		link.click();
		document.body.removeChild(link);
		URL.revokeObjectURL(href);
}

export default Save
